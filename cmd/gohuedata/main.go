package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cheshire137/gohuedata/pkg/bridge_display"
	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/light_loader"
	"github.com/cheshire137/gohuedata/pkg/options"
	"github.com/cheshire137/gohuedata/pkg/sensor_loader"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func main() {
	configPath := "config.yml"
	config, err := config.NewConfig(configPath)
	if err != nil {
		util.LogError("Failed to load configuration:", err)
		return
	}
	util.LogSuccess("Loaded configuration file %s", configPath)

	options := options.ParseOptions()
	bridges := config.Bridges
	bridgeDisplay := bridge_display.NewBridgeDisplay(bridges)
	bridge := bridgeDisplay.GetBridgeSelection(options.BridgeSelection)
	util.LogSuccess("Selected bridge: %s", bridge.Name)

	db, err := sql.Open("sqlite3", config.DatabaseFile)
	if err != nil {
		util.LogError("Failed to open database:", err)
		return
	}
	util.LogSuccess("Loaded %s database", config.DatabaseFile)
	defer db.Close()

	dataStore, err := data_store.NewDataStore(db)
	if err != nil {
		util.LogError("Failed to create tables:", err)
		return
	}

	err = dataStore.AddHueBridge(bridge)
	if err != nil {
		util.LogError("Failed to record Hue bridge:", err)
		return
	}

	bridgeApiUrl, err := bridge.GetApiUrl()
	if err != nil {
		util.LogError("Failed to get bridge URL:", err)
		return
	}

	fahrenheit := options.FahrenheitSpecified(config.FahrenheitSpecified())
	hueClient := hue_api.NewClient(bridgeApiUrl, fahrenheit)

	if options.LoadLights() {
		lightLoader, err := light_loader.NewLightLoader(hueClient)
		if err != nil {
			util.LogError("Failed to load lights:", err)
			return
		}
		lightLoader.DisplayLights()
	}

	if options.LoadSensors() {
		sensorLoader, err := sensor_loader.NewSensorLoader(hueClient, options.SensorSelection)
		if err != nil {
			util.LogError("Failed to load sensors:", err)
			return
		}
		sensorLoader.DisplaySensors()

		tempSensorCount := sensorLoader.TotalTemperatureSensors()
		if tempSensorCount > 0 {
			err = sensorLoader.SaveTemperatureSensorReadings(bridge, dataStore, fahrenheit)
			if err != nil {
				util.LogError("Failed to save temperature readings:", err)
				return
			}
			units := util.Pluralize(tempSensorCount, "reading", "readings")
			util.LogSuccess("Recorded %d temperature %s", tempSensorCount, units)
		}
	}
}
