package main

import (
	"flag"
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

	var bridgeSelection int
	flag.IntVar(&bridgeSelection, "b", 0, "Philips Hue bridge index from config.yml, starts at 1")

	var lightSelection string
	flag.StringVar(&lightSelection, "lights", "all", "Whether to load lights; defaults to all; choose 'all' or 'none'")

	var sensorSelection string
	flag.StringVar(&sensorSelection, "sensors", "all",
		"Which sensors to display, if any; defaults to all; choose 'all', 'motion', 'temperature', or 'none'")

	var tempUnits string
	flag.StringVar(&tempUnits, "t", "", "Temperature units to use; choose from `F` for Fahrenheit or `C` for "+
		"Celsius; defaults to the temperature_units setting in config.yml")

	flag.Parse()

	loadSensors := sensorSelection != "none"
	onlyTempSensors := sensorSelection == "temperature"
	onlyMotionSensors := sensorSelection == "motion"
	loadTempSensors := loadSensors && !onlyMotionSensors
	loadMotionSensors := loadSensors && !onlyTempSensors
	loadAllSensors := loadSensors && !onlyTempSensors && !onlyMotionSensors

	bridges := config.Bridges
	var bridge *hueapi.Bridge
	if isValidBridgeIndex(bridgeSelection, bridges) {
		bridge = &bridges[bridgeSelection-1]
	} else if len(bridges) == 1 {
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

	var fahrenheit bool
	if tempUnits == "" { // not overridden via command line
		fahrenheit = config.TemperatureUnits == "F"
	} else {
		fahrenheit = tempUnits == "F"
	}
	hueClient := hueapi.NewClient(bridgeApiUrl, fahrenheit)

	if lightSelection == "all" {
		lights, err := hueClient.GetLights()
		if err != nil {
			fmt.Println("❌ Failed to get lights:", err)
			return
		}
		fmt.Printf("\n✅ Got %d light(s):\n", len(lights))
		for i, light := range lights {
			fmt.Printf("%d. %s\n", i+1, light.String())
		}
	}

	if loadSensors {
		sensors, err := hueClient.GetSensors(sensorSelection)
		if err != nil {
			fmt.Println("❌ Failed to get sensors:", err)
			return
		}

		if loadAllSensors {
			totalSensors := len(sensors)
			units := util.Pluralize(totalSensors, "sensor", "sensors")
			fmt.Printf("\n✅ Got %d %s:\n", totalSensors, units)
		}

		tempSensors := []*hueapi.TemperatureSensor{}
		motionSensors := []*hueapi.MotionSensor{}
		count := 1

		for _, sensor := range sensors {
			if loadTempSensors {
				tempSensor, ok := sensor.(*hueapi.TemperatureSensor)
				if ok {
					tempSensors = append(tempSensors, tempSensor)
					continue
				}
			}

			if loadMotionSensors {
				motionSensor, ok := sensor.(*hueapi.MotionSensor)
				if ok {
					motionSensors = append(motionSensors, motionSensor)
					continue
				}
			}

			if loadAllSensors {
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
			if loadAllSensors {
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
			if loadAllSensors {
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
