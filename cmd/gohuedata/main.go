package main

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/bridge_display"
	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/hueapi"
	"github.com/cheshire137/gohuedata/pkg/light_display"
	"github.com/cheshire137/gohuedata/pkg/options"
	"github.com/cheshire137/gohuedata/pkg/sensor_loader"
)

func main() {
	config, err := config.NewConfig("config.yml")
	if err != nil {
		fmt.Println("❌ Failed to load configuration:", err)
		return
	}
	fmt.Println("✅ Loaded configuration")

	options := options.ParseOptions()
	bridges := config.Bridges
	bridgeDisplay := bridge_display.NewBridgeDisplay(bridges)
	bridge := bridgeDisplay.GetBridgeSelection(options.BridgeSelection)
	fmt.Println("✅ Selected bridge:", bridge.Name)

	bridgeApiUrl, err := bridge.GetApiUrl()
	if err != nil {
		fmt.Println("❌ Failed to get bridge URL:", err)
		return
	}

	fahrenheit := options.FahrenheitSpecified(config.FahrenheitSpecified())
	hueClient := hueapi.NewClient(bridgeApiUrl, fahrenheit)

	if options.LoadLights() {
		lightDisplay, err := light_display.NewLightDisplay(hueClient)
		if err == nil {
			lightDisplay.DisplayLights()
		} else {
			fmt.Println("❌ Failed to get lights:", err)
		}
	}

	if options.LoadSensors() {
		sensorLoader, err := sensor_loader.NewSensorLoader(hueClient, options.SensorSelection)
		if err == nil {
			sensorLoader.DisplaySensors()
		} else {
			fmt.Println("❌ Failed to get sensors:", err)
		}
	}
}
