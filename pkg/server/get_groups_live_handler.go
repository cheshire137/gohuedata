package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func (e *Env) GetGroupsLiveHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogRequest(r)

	config, err := config.NewConfig(e.options.ConfigPath)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	bridgeName := r.URL.Query().Get("bridge")
	fahrenheit := r.URL.Query().Get("fahrenheit") != "0"

	allBridges := config.Bridges
	var selectedBridges []*hue_api.Bridge
	if bridgeName == "" {
		selectedBridges = allBridges
	} else {
		bridgeName = strings.ToLower(bridgeName)
		for _, bridge := range allBridges {
			if strings.ToLower(bridge.Name) == bridgeName {
				selectedBridges = append(selectedBridges, bridge)
			}
		}
	}

	var groups []*data_store.Group
	totalGroups := 0

	for _, bridge := range selectedBridges {
		bridgeApiUrl, err := bridge.GetApiUrl()
		if err != nil {
			util.ErrorJson(w, err)
			return
		}

		hueClient := hue_api.NewClient(bridgeApiUrl, fahrenheit)
		hueApiGroups, err := hueClient.GetGroups()
		if err != nil {
			util.ErrorJson(w, err)
			return
		}

		bridgeForResponse := &data_store.HueBridge{
			ID:        bridge.IPAddress,
			Name:      bridge.Name,
			IPAddress: bridge.IPAddress,
		}
		totalGroups += len(hueApiGroups)

		for _, hueApiGroup := range hueApiGroups {
			groupForResponse := &data_store.Group{
				ID:           hueApiGroup.ID,
				Name:         hueApiGroup.Name,
				Bridge:       bridgeForResponse,
				Type:         hueApiGroup.Type,
				TotalLights:  len(hueApiGroup.LightIDs),
				TotalSensors: len(hueApiGroup.SensorIDs),
				Class:        hueApiGroup.Class,
			}
			groups = append(groups, groupForResponse)
		}
	}

	response := data_store.GroupsLiveResponse{
		Groups:     groups,
		Page:       1,
		TotalPages: 1,
		TotalCount: totalGroups,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}