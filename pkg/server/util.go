package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cheshire137/gohuedata/pkg/config"
	"github.com/cheshire137/gohuedata/pkg/data_store"
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

func ErrorJson(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	if err == sql.ErrNoRows {
		statusCode = http.StatusNotFound
	} else {
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	response := data_store.ErrorResponse{Error: err.Error()}
	json.NewEncoder(w).Encode(response)
}
