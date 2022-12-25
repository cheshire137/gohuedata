package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/sensor_loader"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type TemperatureSensorsLiveResponse struct {
	TemperatureSensors []*data_store.TemperatureSensorExtended `json:"temperatureSensors"`
	Page               int                                     `json:"page"`
	TotalPages         int                                     `json:"totalPages"`
	TotalCount         int                                     `json:"totalCount"`
}

func (e *Env) GetTemperatureSensorsLiveHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogInfo("GET %s", r.URL.Path)

	config, err := config.NewConfig(e.options.ConfigPath)
	if err != nil {
		util.LogError("Failed to load configuration:", err)
		return
	}

	bridgeName := r.URL.Query().Get("bridge")
	fahrenheit := r.URL.Query().Get("fahrenheit") != "0"

	allBridges := config.Bridges
	var selectedBridges []*hue_api.Bridge
	if bridgeName == "" {
		selectedBridges = allBridges
	} else {
		bridgeName = strings.ToLower(bridgeName)
		for _, bridge := range allBridges {
			if strings.ToLower(bridge.Name) == bridgeName {
				selectedBridges = append(selectedBridges, bridge)
			}
		}
	}

	var tempSensors []*data_store.TemperatureSensorExtended
	totalTempSensors := 0

	for _, bridge := range selectedBridges {
		bridgeApiUrl, err := bridge.GetApiUrl()
		if err != nil {
			util.ErrorJson(w, err)
			return
		}

		hueClient := hue_api.NewClient(bridgeApiUrl, fahrenheit)
		sensorLoader, err := sensor_loader.NewSensorLoader(hueClient, "temperature")
		if err != nil {
			util.ErrorJson(w, err)
			return
		}

		bridgeForResponse := &data_store.HueBridge{
			ID:        bridge.IPAddress,
			Name:      bridge.Name,
			IPAddress: bridge.IPAddress,
		}
		tempSensorsForBridge := sensorLoader.TemperatureSensors
		totalTempSensors += sensorLoader.TotalTemperatureSensors()

		for _, tempSensor := range tempSensorsForBridge {
			temperature := tempSensor.Temperature()
			units := tempSensor.TempUnits()
			tempSensorForResponse := data_store.TemperatureSensor{
				ID:          tempSensor.UniqueID,
				Name:        tempSensor.Name,
				Bridge:      bridgeForResponse,
				LastUpdated: tempSensor.State.LastUpdated,
			}
			readingForResponse := &data_store.TemperatureReading{
				ID:                fmt.Sprintf("%s%s%.1f%s", tempSensor.UniqueID, tempSensor.State.LastUpdated, temperature, units),
				TemperatureSensor: &tempSensorForResponse,
				Timestamp:         tempSensor.State.LastUpdated,
				Temperature:       temperature,
				Units:             units,
			}
			extendedTempSensor := &data_store.TemperatureSensorExtended{
				TemperatureSensor: tempSensorForResponse,
				LatestReading:     readingForResponse,
			}
			tempSensors = append(tempSensors, extendedTempSensor)
		}
	}

	response := TemperatureSensorsLiveResponse{
		TemperatureSensors: tempSensors,
		Page:               1,
		TotalPages:         1,
		TotalCount:         totalTempSensors,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
