package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/pagination"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func (e *Env) GetDailyTemperatureReadingsHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogRequest(r)

	pageInfo, err := pagination.GetPageInfoParams(r.URL)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	bridgeName := r.URL.Query().Get("bridge")
	updatedSince := r.URL.Query().Get("updated_since")
	updatedBefore := r.URL.Query().Get("updated_before")
	sensorID := r.URL.Query().Get("sensor_id")
	fahrenheit := r.URL.Query().Get("fahrenheit") != "0"

	filter := &data_store.TemperatureReadingFilter{
		Page:          pageInfo.Page,
		PerPage:       pageInfo.PerPage,
		BridgeName:    bridgeName,
		UpdatedSince:  updatedSince,
		UpdatedBefore: updatedBefore,
		SensorID:      sensorID,
		Fahrenheit:    fahrenheit,
	}

	tempReadingSummaries, err := e.ds.LoadDailyTemperatureReadings(filter)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	totalDailyTempReadings, err := e.ds.TotalDailyTemperatureReadings(filter)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	response := data_store.TemperatureReadingSummariesResponse{
		TemperatureReadingSummaries: tempReadingSummaries,
		Page:                        pageInfo.Page,
		TotalPages:                  util.TotalPages(totalDailyTempReadings, pageInfo.PerPage),
		TotalCount:                  totalDailyTempReadings,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
