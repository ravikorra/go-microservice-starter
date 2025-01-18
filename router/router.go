package router

import (
	authenticationservice "go-microservice-starter/services/authentication_service_handle"

	"github.com/gorilla/mux"
)

// InitRouter initializes and returns the mux router
func InitRouter() *mux.Router {
	r := mux.NewRouter()

	// Apply global middleware
	//r.Use(middleware.LoggingMiddleware)

	// Register routes
	r.HandleFunc("/home", authenticationservice.HomeHandler).Methods("GET")
	//r.HandleFunc("/secure", handlers.SecureHandler).Methods("GET")
	//r.HandleFunc("/", ).Methods("GET")

	return r
}
