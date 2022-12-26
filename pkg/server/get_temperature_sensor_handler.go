package server

import (
	"encoding/json"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/util"
)

func (e *Env) GetTemperatureSensorHandler(w http.ResponseWriter, r *http.Request) {
	e.enableCors(&w)
	util.LogRequest(r)

	sensorID := r.URL.Query().Get("id")
	tempSensor, err := e.ds.LoadTemperatureSensor(sensorID)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempSensor)
}
