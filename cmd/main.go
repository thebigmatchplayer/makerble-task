package main

import (
	"net/http"

	"github.com/thebigmatchplayer/markerble-task/config"
	"github.com/thebigmatchplayer/markerble-task/handlers"
	"github.com/thebigmatchplayer/markerble-task/middleware"
	"go.uber.org/zap"
)

func main() {
	config.InitLogger()
	config.InitDB()
	defer config.DB.Close()

	if err := middleware.LoadAccessMatrix("middleware/access.yaml"); err != nil {
		config.Log.Fatal("Failed to load access matrix", zap.Error(err))
	}

	handlers.SetupRoutes()

	config.Log.Info("Starting CRUD Portal on :9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		config.Log.Fatal("Server failed", zap.Error(err))
	}
}
