package server

import (
	"strings"

	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

func GetSelectedBridges(bridgeName string, config *config.Config) []*hue_api.Bridge {
	allBridges := config.Bridges

	if bridgeName == "" {
		return allBridges
	}

	var selectedBridges []*hue_api.Bridge
	bridgeName = strings.ToLower(bridgeName)

	for _, bridge := range allBridges {
		if strings.ToLower(bridge.Name) == bridgeName {
			selectedBridges = append(selectedBridges, bridge)
		}
	}

	return selectedBridges
}
