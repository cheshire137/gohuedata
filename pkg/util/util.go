package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/cheshire137/gohuedata/pkg/data_store"
)

func Pluralize(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

func LogInfo(format string, a ...interface{}) {
	fmt.Printf("ℹ️ "+format+"\n", a...)
}

func LogSuccess(format string, a ...interface{}) {
	fmt.Printf("✅ "+format+"\n", a...)
}

func LogError(a ...interface{}) {
	fmt.Print("❌ ")
	fmt.Println(a...)
}

func ErrorJson(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	if err == sql.ErrNoRows {
		statusCode = http.StatusNotFound
	} else {
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	response := data_store.ErrorResponse{Error: err.Error()}
	json.NewEncoder(w).Encode(response)
}

func TotalPages(count int, perPage int) int {
	return int(math.Ceil(float64(count) / float64(perPage)))
}

func LinePunctuation(quietMode bool) string {
	if quietMode {
		return ""
	}
	return ":"
}

func LogRequest(r *http.Request) {
	queryStr := r.URL.RawQuery
	if queryStr == "" {
		LogInfo("%s %s", r.Method, r.URL.Path)
	} else {
		LogInfo("%s %s?%s", r.Method, r.URL.Path, r.URL.RawQuery)
	}
}
