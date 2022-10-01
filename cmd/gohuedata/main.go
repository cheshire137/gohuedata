package main

import (
	"fmt"
	"strconv"

	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/hueapi"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func displayBridgePrompt(bridges []hueapi.Bridge) {
	fmt.Println("\nChoose a Philips Hue bridge:")
	for i, bridge := range bridges {
		fmt.Printf("%d. %s\n", i+1, bridge.String())
	}
}

func isValidBridgeIndex(index int, bridges []hueapi.Bridge) bool {
	return index >= 1 && index <= len(bridges)
}

func getBridgeSelectionFromUser(bridges []hueapi.Bridge) *hueapi.Bridge {
	displayBridgePrompt(bridges)

	userInput := ""
	bridgeSelection := -1
	invalidIndex := true
	errorPrefix := "❌ Invalid bridge choice"
	var err error

	for userInput == "" || invalidIndex {
		fmt.Scanln(&userInput)
		bridgeSelection, err = strconv.Atoi(userInput)
		nonIntegerInput := err != nil
		invalidIndex = !isValidBridgeIndex(bridgeSelection, bridges)
		if nonIntegerInput {
			fmt.Println(errorPrefix + ":" + err.Error())
		} else if invalidIndex {
			if len(bridges) > 1 {
				fmt.Println(errorPrefix+", must be between 1 and", len(bridges))
			} else {
				fmt.Println(errorPrefix + ", must be 1")
			}
		}
		if nonIntegerInput || invalidIndex {
			displayBridgePrompt(bridges)
			userInput = ""
		}
	}

	bridgeIndex := bridgeSelection - 1
	return &bridges[bridgeIndex]
}

func main() {
	config, err := config.NewConfig("config.yml")
	if err != nil {
		fmt.Println("❌ Failed to load configuration:", err)
		return
	}
	fmt.Println("✅ Loaded configuration")

	bridges := config.Bridges
	var bridge *hueapi.Bridge
	if len(bridges) == 1 {
		bridge = &bridges[0]
	} else {
		bridge = getBridgeSelectionFromUser(config.Bridges)
	}
	fmt.Println("✅ Selected bridge:", bridge.Name)

	bridgeApiUrl, err := bridge.GetApiUrl()
	if err != nil {
		fmt.Println("❌ Failed to get bridge URL:", err)
		return
	}

	hueClient := hueapi.NewClient(bridgeApiUrl)

	lights, err := hueClient.GetLights()
	if err != nil {
		fmt.Println("❌ Failed to get lights:", err)
		return
	}
	fmt.Printf("\n✅ Got %d light(s):\n", len(lights))
	for i, light := range lights {
		fmt.Printf("%d. %s\n", i+1, light.String())
	}

	sensors, err := hueClient.GetSensors()
	if err != nil {
		fmt.Println("❌ Failed to get sensors:", err)
		return
	}
	totalSensors := len(sensors)
	units := util.Pluralize(totalSensors, "sensor", "sensors")
	fmt.Printf("\n✅ Got %d %s:\n", totalSensors, units)

	tempSensors := []*hueapi.TemperatureSensor{}
	motionSensors := []*hueapi.MotionSensor{}
	count := 1

	for _, sensor := range sensors {
		tempSensor, ok := sensor.(*hueapi.TemperatureSensor)
		if ok {
			tempSensors = append(tempSensors, tempSensor)
			continue
		}

		motionSensor, ok := sensor.(*hueapi.MotionSensor)
		if ok {
			motionSensors = append(motionSensors, motionSensor)
			continue
		}

		sensor, ok := sensor.(*hueapi.Sensor)
		if !ok {
			fmt.Println("❌ Unknown sensor type:", sensor)
			continue
		}
		fmt.Printf("%d. %s\n", count, sensor.String())
		count++
	}

	totalTempSensors := len(tempSensors)
	if totalTempSensors > 0 {
		units := util.Pluralize(totalTempSensors, "sensor", "sensors")
		fmt.Printf("\n✅ Including %d temperature %s:\n", totalTempSensors, units)
		for _, sensor := range tempSensors {
			fmt.Printf("%d. %s\n", count, sensor.String())
			count++
		}
	}

	totalMotionSensors := len(motionSensors)
	if totalMotionSensors > 0 {
		units := util.Pluralize(totalMotionSensors, "sensor", "sensors")
		fmt.Printf("\n✅ Including %d motion %s:\n", totalMotionSensors, units)
		for _, sensor := range motionSensors {
			fmt.Printf("%d. %s\n", count, sensor.String())
			count++
		}
	}
}
