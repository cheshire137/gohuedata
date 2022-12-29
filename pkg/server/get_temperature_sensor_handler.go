package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func (e *Env) GetTemperatureSensorHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogRequest(r)

	sensorID := r.URL.Query().Get("id")
	fahrenheit := r.URL.Query().Get("fahrenheit") != "0"

	tempSensor, err := e.ds.LoadTemperatureSensor(sensorID, fahrenheit)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	maxTemp, err := e.ds.LoadMaxRecordedTemperatureForSensor(sensorID, fahrenheit)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	minTemp, err := e.ds.LoadMinRecordedTemperatureForSensor(sensorID, fahrenheit)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	avgTemp, err := e.ds.LoadAvgRecordedTemperatureForSensor(sensorID, fahrenheit)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	response := data_store.TemperatureSensorResponse{
		TemperatureSensor: tempSensor,
		MaxTemperature:    maxTemp,
		MinTemperature:    minTemp,
		AvgTemperature:    avgTemp,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
