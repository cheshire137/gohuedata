package bridge_display

import (
	"fmt"
	"strconv"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

type BridgeDisplay struct {
	Bridges []hue_api.Bridge
}

func NewBridgeDisplay(bridges []hue_api.Bridge) *BridgeDisplay {
	return &BridgeDisplay{Bridges: bridges}
}

func (bd *BridgeDisplay) GetBridgeSelection(bridgeSelection int) *hue_api.Bridge {
	if bd.IsValidBridgeIndex(bridgeSelection) {
		return &bd.Bridges[bridgeSelection-1]
	}
	if len(bd.Bridges) == 1 {
		return &bd.Bridges[0]
	}
	return bd.GetBridgeSelectionFromUser()
}

func (bd *BridgeDisplay) GetBridgeSelectionFromUser() *hue_api.Bridge {
	bd.DisplayBridgePrompt()

	userInput := ""
	bridgeSelection := -1
	invalidIndex := true
	errorPrefix := "❌ Invalid bridge choice"
	var err error

	for userInput == "" || invalidIndex {
		fmt.Scanln(&userInput)
		bridgeSelection, err = strconv.Atoi(userInput)
		nonIntegerInput := err != nil
		invalidIndex = !bd.IsValidBridgeIndex(bridgeSelection)
		if nonIntegerInput {
			fmt.Println(errorPrefix + ":" + err.Error())
		} else if invalidIndex {
			if len(bd.Bridges) > 1 {
				fmt.Println(errorPrefix+", must be between 1 and", len(bd.Bridges))
			} else {
				fmt.Println(errorPrefix + ", must be 1")
			}
		}
		if nonIntegerInput || invalidIndex {
			bd.DisplayBridgePrompt()
			userInput = ""
		}
	}

	bridgeIndex := bridgeSelection - 1
	return &bd.Bridges[bridgeIndex]
}

func (bd *BridgeDisplay) DisplayBridgePrompt() {
	fmt.Println("\nChoose a Philips Hue bridge:")
	for i, bridge := range bd.Bridges {
		fmt.Printf("%d. %s\n", i+1, bridge.String())
	}
}

func (bd *BridgeDisplay) IsValidBridgeIndex(index int) bool {
	return index >= 1 && index <= len(bd.Bridges)
}
