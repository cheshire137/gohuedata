package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	var lightsInGroup []*data_store.LightExtended

	if totalLights > 0 {
		lightsInGroup = make([]*data_store.LightExtended, totalLights)
		lightLoader, err := light_loader.NewLightLoader(hueClient)
		if err != nil {
			ErrorJson(w, err)
			return
		}
		timestamp := time.Now().Format(time.RFC3339)
		for i, lightID := range hueGroup.LightIDs {
			hueLight, ok := lightLoader.LightsByID[lightID]
			if ok {
				light := data_store.Light{
					UniqueID: hueLight.UniqueID,
					ID:       lightID,
					Name:     hueLight.Name,
					Bridge:   bridgeForResponse,
				}
				lightsInGroup[i] = &data_store.LightExtended{
					Light: light,
					LatestState: &data_store.LightState{
						Timestamp: timestamp,
						On:        hueLight.State.On,
						Light:     &light,
					},
				}
			}
		}
	}

	groupExtended := data_store.GroupExtended{
		ID:           groupID,
		UniqueID:     fmt.Sprintf("%s-%s", bridge.IPAddress, hueGroup.ID),
		Name:         hueGroup.Name,
		Type:         hueGroup.Type,
		Class:        hueGroup.Class,
		Bridge:       bridgeForResponse,
		TotalLights:  totalLights,
		TotalSensors: len(hueGroup.SensorIDs),
		Lights:       lightsInGroup,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groupExtended)
}
