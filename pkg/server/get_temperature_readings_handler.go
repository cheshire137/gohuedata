package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/pagination"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type TemperatureReadingsResponse struct {
	TemperatureReadings []*data_store.TemperatureReading `json:"temperatureReadings"`
	Page                int                              `json:"page"`
	TotalPages          int                              `json:"totalPages"`
	TotalCount          int                              `json:"totalCount"`
}

func (e *Env) GetTemperatureReadingsHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogRequest(r)

	pageInfo, err := pagination.GetPageInfoParams(r.URL)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	bridgeName := r.URL.Query().Get("bridge")
	updatedSince := r.URL.Query().Get("updated_since")
	updatedBefore := r.URL.Query().Get("updated_before")
	sensorID := r.URL.Query().Get("sensor_id")
	filter := &data_store.TemperatureReadingFilter{
		Page:          pageInfo.Page,
		PerPage:       pageInfo.PerPage,
		BridgeName:    bridgeName,
		UpdatedSince:  updatedSince,
		UpdatedBefore: updatedBefore,
		SensorID:      sensorID,
	}

	tempReadings, err := e.ds.LoadTemperatureReadings(filter)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	totalTempReadings, err := e.ds.TotalTemperatureReadings(filter)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	response := TemperatureReadingsResponse{
		TemperatureReadings: tempReadings,
		Page:                pageInfo.Page,
		TotalPages:          util.TotalPages(totalTempReadings, pageInfo.PerPage),
		TotalCount:          totalTempReadings,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
