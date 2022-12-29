package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/light_loader"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func (e *Env) GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogRequest(r)

	config, err := config.NewConfig(e.options.ConfigPath)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	groupID := r.URL.Query().Get("id")
	if groupID == "" {
		ErrorMessageJson(w, "`id` parameter is required, specifying Philips Hue group ID")
		return
	}

	bridgeName := r.URL.Query().Get("bridge")
	if bridgeName == "" {
		ErrorMessageJson(w, "`bridge` parameter is required, specifying name of Philips Hue bridge that group is on")
		return
	}

	selectedBridges := GetSelectedBridges(bridgeName, config)
	if len(selectedBridges) != 1 {
		ErrorMessageJson(w, "`bridge` parameter didn't match a known Philips Hue bridge")
		return
	}

	bridge := selectedBridges[0]
	var groupForResponse data_store.Group

	bridgeApiUrl, err := bridge.GetApiUrl()
	if err != nil {
		ErrorJson(w, err)
		return
	}

	hueClient := hue_api.NewClient(bridgeApiUrl, true)
	hueGroup, err := hueClient.GetGroup(groupID)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	bridgeForResponse := &data_store.HueBridge{
		ID:        bridge.IPAddress,
		Name:      bridge.Name,
		IPAddress: bridge.IPAddress,
	}

	totalLights := len(hueGroup.LightIDs)
	var lightsInGroup []*data_store.Light

	if totalLights > 0 {
		lightsInGroup = make([]*data_store.Light, totalLights)
		lightLoader, err := light_loader.NewLightLoader(hueClient)
		if err != nil {
			ErrorJson(w, err)
			return
		}
		for i, lightID := range hueGroup.LightIDs {
			light, ok := lightLoader.LightsByID[lightID]
			if ok {
				lightsInGroup[i] = &data_store.Light{
					UniqueID: light.UniqueID,
					ID:       lightID,
					Name:     light.Name,
					Bridge:   bridgeForResponse,
				}
			}
		}
	}

	groupForResponse = data_store.Group{
		ID:           groupID,
		Name:         hueGroup.Name,
		Type:         hueGroup.Type,
		Class:        hueGroup.Class,
		Bridge:       bridgeForResponse,
		TotalLights:  totalLights,
		TotalSensors: len(hueGroup.SensorIDs),
		Lights:       lightsInGroup,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groupForResponse)
}
