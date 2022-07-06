package main

import (
	"fmt"
	"strconv"

	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/hueapi"
)

func displayBridgePrompt(bridges []hueapi.Bridge) {
	fmt.Println("\nChoose a Philips Hue bridge:")
	for i, bridge := range bridges {
		fmt.Printf("%d. %s -- %s via %s\n", i+1, bridge.Name, bridge.IPAddress, bridge.Username)
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

	bridgeUrl, err := bridge.GetUrl()
	if err != nil {
		fmt.Println("❌ Failed to get bridge URL:", err)
		return
	}

	hueClient := hueapi.NewClient(bridgeUrl.String(), bridge.Username)
	lights, err := hueClient.GetLights()
	if err != nil {
		fmt.Println("❌ Failed to get lights:", err)
		return
	}
	fmt.Printf("✅ Got %d lights:\n", len(lights))
	for i, light := range lights {
		fmt.Printf("%d. %s\n", i+1, light.Name)
	}
}
