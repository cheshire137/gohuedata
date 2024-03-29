package bridge_display

import (
	"fmt"
	"strconv"

	"github.com/cheshire137/gohuedata/pkg/cli_options"
	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

type BridgeDisplay struct {
	Bridges []*hue_api.Bridge
}

func NewBridgeDisplay(bridges []*hue_api.Bridge) *BridgeDisplay {
	return &BridgeDisplay{Bridges: bridges}
}

func (bd *BridgeDisplay) GetBridgeSelection(bridgeSelection int) []*hue_api.Bridge {
	if cli_options.AllBridges == bridgeSelection {
		return bd.Bridges
	}
	if bd.IsValidBridgeIndex(bridgeSelection) {
		index := bridgeSelection - 1
		return bd.Bridges[index : index+1]
	}
	if len(bd.Bridges) == 1 {
		return bd.Bridges[0:1]
	}
	return bd.GetBridgeSelectionFromUser()
}

func (bd *BridgeDisplay) GetBridgeSelectionFromUser() []*hue_api.Bridge {
	bd.DisplayBridgePrompt()

	userInput := ""
	bridgeSelection := 0
	invalidSelection := true
	errorPrefix := "❌ Invalid bridge choice"
	var err error

	for userInput == "" || invalidSelection {
		fmt.Scanln(&userInput)
		bridgeSelection, err = strconv.Atoi(userInput)
		nonIntegerInput := err != nil
		invalidSelection = !bd.IsValidBridgeIndex(bridgeSelection) && bridgeSelection != cli_options.AllBridges
		if nonIntegerInput {
			fmt.Println(errorPrefix + ":" + err.Error())
		} else if invalidSelection {
			if len(bd.Bridges) > 1 {
				fmt.Println(errorPrefix+", must be between 1 and", len(bd.Bridges))
			} else {
				fmt.Println(errorPrefix + ", must be 1")
			}
		}
		if nonIntegerInput || invalidSelection {
			bd.DisplayBridgePrompt()
			userInput = ""
		}
	}

	if bridgeSelection == cli_options.AllBridges {
		return bd.Bridges
	}

	bridgeIndex := bridgeSelection - 1
	return bd.Bridges[bridgeIndex : bridgeIndex+1]
}

func (bd *BridgeDisplay) DisplayBridgePrompt() {
	fmt.Println("\nChoose a Philips Hue bridge:")
	fmt.Printf("%d. All bridges\n", cli_options.AllBridges)
	for i, bridge := range bd.Bridges {
		fmt.Printf("%d. %s\n", i+1, bridge.String())
	}
}

func (bd *BridgeDisplay) IsValidBridgeIndex(index int) bool {
	return index >= 1 && index <= len(bd.Bridges)
}
