package main

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/bridge_display"
	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/hueapi"
	"github.com/cheshire137/gohuedata/pkg/light_display"
	"github.com/cheshire137/gohuedata/pkg/options"
	"github.com/cheshire137/gohuedata/pkg/util"
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
		lightDisplay := light_display.NewLightDisplay(hueClient)
		lightDisplay.LoadLights()
	}

	if options.LoadSensors() {
		sensors, err := hueClient.GetSensors(options.SensorSelection)
		if err != nil {
			fmt.Println("❌ Failed to get sensors:", err)
			return
		}

		if options.LoadAllSensors() {
			totalSensors := len(sensors)
			units := util.Pluralize(totalSensors, "sensor", "sensors")
			fmt.Printf("\n✅ Got %d %s:\n", totalSensors, units)
		}

		tempSensors := []*hueapi.TemperatureSensor{}
		motionSensors := []*hueapi.MotionSensor{}
		count := 1

		for _, sensor := range sensors {
			if options.LoadTemperatureSensors() {
				tempSensor, ok := sensor.(*hueapi.TemperatureSensor)
				if ok {
					tempSensors = append(tempSensors, tempSensor)
					continue
				}
			}

			if options.LoadMotionSensors() {
				motionSensor, ok := sensor.(*hueapi.MotionSensor)
				if ok {
					motionSensors = append(motionSensors, motionSensor)
					continue
				}
			}

			if options.LoadAllSensors() {
				sensor, ok := sensor.(*hueapi.Sensor)
				if !ok {
					fmt.Println("❌ Unknown sensor type:", sensor)
					continue
				}
				fmt.Printf("%d. %s\n", count, sensor.String())
				count++
			}
		}

		totalTempSensors := len(tempSensors)
		if totalTempSensors > 0 {
			units := util.Pluralize(totalTempSensors, "sensor", "sensors")
			var intro string
			if options.LoadAllSensors() {
				intro = "Including"
			} else {
				intro = "Got"
			}
			fmt.Printf("\n✅ %s %d temperature %s:\n", intro, totalTempSensors, units)
			for _, sensor := range tempSensors {
				fmt.Printf("%d. %s\n", count, sensor.String())
				count++
			}
		}

		totalMotionSensors := len(motionSensors)
		if totalMotionSensors > 0 {
			units := util.Pluralize(totalMotionSensors, "sensor", "sensors")
			var intro string
			if options.LoadAllSensors() {
				intro = "Including"
			} else {
				intro = "Got"
			}
			fmt.Printf("\n✅ %s %d motion %s:\n", intro, totalMotionSensors, units)
			for _, sensor := range motionSensors {
				fmt.Printf("%d. %s\n", count, sensor.String())
				count++
			}
		}
	}
}
