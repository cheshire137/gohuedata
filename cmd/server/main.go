package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/server"
	options "github.com/cheshire137/gohuedata/pkg/server_options"
	"github.com/cheshire137/gohuedata/pkg/util"
)

func main() {
	options := options.ParseOptions()

	db, err := sql.Open("sqlite3", options.DatabasePath)
	if err != nil {
		util.LogError("Failed to open database:", err)
		return
	}
	util.LogSuccess("Loaded %s database", options.DatabasePath)
	defer db.Close()

	dataStore := data_store.NewDataStore(db)
	mux := http.NewServeMux()
	env := server.NewEnv(dataStore, options)

	mux.Handle("/", http.HandlerFunc(env.RootHandler))
	mux.Handle("/api/temperature-readings", http.HandlerFunc(env.GetTemperatureReadingsHandler))

	server := &http.Server{
		Addr:    options.Address(),
		Handler: mux,
	}

	util.LogInfo("Starting server at http://localhost:%d", options.Port)
	go func(srv *http.Server) {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			util.LogError("Could not start server:", err)
		}
	}(server)

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
	<-stopSignal

	shutdownServer(server)
}

func shutdownServer(server *http.Server) {
	util.LogInfo("Stopping server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		util.LogError("Could not cleanly stop server:", err)
	}
}
