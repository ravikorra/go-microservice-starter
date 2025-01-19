package main

import (
	"go-microservice-starter/log"
	"go-microservice-starter/router"
	"net/http"

	"go.uber.org/zap"
)

func main() {

	// Initialize the logger
	// Initialize the logger
	if err := log.Initialize(); err != nil {
		panic(err)
	}

	log.Info("Application starting...")

	defer log.Sync() // Ensure logs are flushed before the program exits

	r := router.InitRouter()

	log.Info("Server is starting", zap.String("url", "http://localhost:4100"))

	// Start the server
	if err := http.ListenAndServe(":4100", r); err != nil {
		log.Error("Server failed to start", zap.Error(err))
		log.Sync() // Ensure logs are flushed before exiting
	}

}
