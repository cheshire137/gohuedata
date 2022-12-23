package server

import "net/http"

func errorJson(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
