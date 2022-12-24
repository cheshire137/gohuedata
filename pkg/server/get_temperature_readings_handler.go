package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type TemperatureReadingsResponse struct {
	TemperatureReadings []*data_store.TemperatureReading `json:"temperatureReadings"`
	Page                int                              `json:"page"`
	TotalPages          int                              `json:"totalPages"`
}

func (e *Env) GetTemperatureReadingsHandler(w http.ResponseWriter, r *http.Request) {
	pageInfo, err := GetPageInfo(r.URL)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	tempReadings, err := e.ds.LoadTemperatureReadings(pageInfo.Page, pageInfo.PerPage)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	totalTempReadings, err := e.ds.TotalTemperatureReadings()
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	response := TemperatureReadingsResponse{
		TemperatureReadings: tempReadings,
		Page:                pageInfo.Page,
		TotalPages:          util.TotalPages(totalTempReadings, pageInfo.PerPage),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
