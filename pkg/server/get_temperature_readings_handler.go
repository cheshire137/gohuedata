package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/util"
)

func (e *Env) GetTemperatureReadingsHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	tempReadings, err := e.ds.LoadTemperatureReadings(page)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempReadings)
}
