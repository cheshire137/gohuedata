package server

import (
	"fmt"
	"net/http"
)

func (e *Env) enableCors(w *http.ResponseWriter) {
	allowedUrl := fmt.Sprintf("http://localhost:%d", e.options.FrontendPort)
	(*w).Header().Set("Access-Control-Allow-Origin", allowedUrl)
}
