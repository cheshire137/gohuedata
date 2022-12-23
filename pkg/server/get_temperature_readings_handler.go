package server

import (
	"encoding/json"
	"net/http"
)

func (e *Env) GetTemperatureReadingsHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	tempReadings, err := e.ds.LoadTemperatureReadings(page)
	if err != nil {
		errorJson(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempReadings)
}
