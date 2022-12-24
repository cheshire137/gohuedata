package util

import (
	"fmt"
	"math"
	"net/http"
)

func Pluralize(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

func LogInfo(format string, a ...interface{}) {
	fmt.Printf("\nℹ️ "+format+"\n", a...)
}

func LogSuccess(format string, a ...interface{}) {
	fmt.Printf("\n✅ "+format+"\n", a...)
}

func LogError(a ...interface{}) {
	fmt.Print("❌ ")
	fmt.Println(a...)
}

func ErrorJson(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func TotalPages(count int, perPage int) int {
	return int(math.Ceil(float64(count) / float64(perPage)))
}
