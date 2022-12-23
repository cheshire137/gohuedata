package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cheshire137/gohuedata/pkg/bridge_display"
	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/hueapi"
	"github.com/cheshire137/gohuedata/pkg/light_loader"
	"github.com/cheshire137/gohuedata/pkg/options"
	"github.com/cheshire137/gohuedata/pkg/sensor_loader"
	"github.com/cheshire137/gohuedata/pkg/temperature_sensor_store"
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

	bridgeApiUrl, err := bridge.GetApiUrl()
	if err != nil {
		util.LogError("Failed to get bridge URL:", err)
		return
	}

	fahrenheit := options.FahrenheitSpecified(config.FahrenheitSpecified())
	hueClient := hueapi.NewClient(bridgeApiUrl, fahrenheit)
	db, err := sql.Open("sqlite3", config.DatabaseFile)
	if err != nil {
		util.LogError("Failed to open database:", err)
		return
	}
	util.LogSuccess("Loaded %s database", config.DatabaseFile)
	defer db.Close()

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

		_, err = temperature_sensor_store.NewTemperatureSensorStore(db)
		if err != nil {
			util.LogError("Failed to create temperature sensor table:", err)
			return
		}
	}
}
