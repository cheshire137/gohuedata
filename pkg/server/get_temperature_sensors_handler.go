package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/pagination"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type TemperatureSensorsResponse struct {
	TemperatureSensors []*data_store.TemperatureSensor `json:"temperatureSensors"`
	Page               int                             `json:"page"`
	TotalPages         int                             `json:"totalPages"`
	TotalCount         int                             `json:"totalCount"`
}

func (e *Env) GetTemperatureSensorsHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogInfo("GET %s", r.URL.Path)

	pageInfo, err := pagination.GetPageInfoParams(r.URL)
	if err != nil {
		util.ErrorJson(w, err)
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
		util.ErrorJson(w, err)
		return
	}

	totalTempSensors, err := e.ds.TotalTemperatureSensors(filter)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	response := TemperatureSensorsResponse{
		TemperatureSensors: tempSensors,
		Page:               pageInfo.Page,
		TotalPages:         util.TotalPages(totalTempSensors, pageInfo.PerPage),
		TotalCount:         totalTempSensors,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}