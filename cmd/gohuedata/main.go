package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cheshire137/gohuedata/pkg/bridge_display"
	options "github.com/cheshire137/gohuedata/pkg/cli_options"
	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/light_loader"
	"github.com/cheshire137/gohuedata/pkg/sensor_loader"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func main() {
	options := options.ParseOptions()
	config, err := config.NewConfig(options.ConfigPath)
	if err != nil {
		util.LogError("Failed to load configuration:", err)
		return
	}

	if !options.QuietMode {
		util.LogSuccess("Loaded configuration file %s", options.ConfigPath)
	}

	allBridges := config.Bridges
	bridgeDisplay := bridge_display.NewBridgeDisplay(allBridges)
	var selectedBridges []*hue_api.Bridge
	if options.LoadAllBridges() {
		selectedBridges = allBridges
	} else {
		selectedBridges = bridgeDisplay.GetBridgeSelection(options.BridgeSelection)
	}

	if !options.QuietMode {
		units := util.Pluralize(len(selectedBridges), "bridge", "bridges")
		bridgeNames := hue_api.BridgeNames(selectedBridges)
		util.LogSuccess("Selected %s: %s", units, bridgeNames)
	}

	db, err := sql.Open("sqlite3", config.DatabaseFile)
	if err != nil {
		util.LogError("Failed to open database:", err)
		return
	}
	defer db.Close()

	if !options.QuietMode {
		util.LogSuccess("Loaded %s database", config.DatabaseFile)
	}

	dataStore := data_store.NewDataStore(db)
	err = dataStore.CreateTables()
	if err != nil {
		util.LogError("Failed to create tables:", err)
		return
	}

	for _, bridge := range selectedBridges {
		if options.LoadAllBridges() {
			util.LogInfo("Bridge %s", bridge.Name)
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

			err = lightLoader.DisplayLights(options.QuietMode)
			if err != nil {
				util.LogError("Failed to display lights:", err)
				return
			}
		}

		if options.LoadSensors() {
			sensorLoader, err := sensor_loader.NewSensorLoader(hueClient, options.SensorSelection)
			if err != nil {
				util.LogError("Failed to load sensors:", err)
				return
			}

			sensorLoader.DisplaySensors(options.QuietMode)

			tempSensorCount := sensorLoader.TotalTemperatureSensors()
			if tempSensorCount > 0 {
				err = sensorLoader.SaveTemperatureSensorReadings(bridge, dataStore, fahrenheit)
				if err != nil {
					util.LogError("Failed to save temperature readings:", err)
					return
				}
				units := util.Pluralize(tempSensorCount, "reading", "readings")
				util.LogSuccess("Recorded %d temperature %s", tempSensorCount, units)
				mostRecentSensorState := sensorLoader.MostRecentlyUpdatedSensorState()
				if mostRecentSensorState != nil {
					util.LogInfo("Most recent reading: %s", mostRecentSensorState.LastUpdatedSummary())
				}
			}
		}
	}
}
