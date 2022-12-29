package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/pagination"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func (e *Env) GetTemperatureSensorsHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogRequest(r)

	pageInfo, err := pagination.GetPageInfoParams(r.URL)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	bridgeName := r.URL.Query().Get("bridge")
	filter := &data_store.TemperatureSensorFilter{
		Page:       pageInfo.Page,
		PerPage:    pageInfo.PerPage,
		BridgeName: bridgeName,
	}

	tempSensors, err := e.ds.LoadTemperatureSensors(filter)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	totalTempSensors, err := e.ds.TotalTemperatureSensors(filter)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	response := data_store.TemperatureSensorsResponse{
		TemperatureSensors: tempSensors,
		Page:               pageInfo.Page,
		TotalPages:         util.TotalPages(totalTempSensors, pageInfo.PerPage),
		TotalCount:         totalTempSensors,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
