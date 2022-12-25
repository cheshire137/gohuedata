package server

import (
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/util"
)

func (e *Env) RootHandler(w http.ResponseWriter, r *http.Request) {
	util.LogInfo("GET %s", r.URL.Path)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
