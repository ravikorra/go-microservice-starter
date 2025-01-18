package main

import (
	"go-microservice-starter/router"
	"log"
	"net/http"
)

func main() {
	r := router.InitRouter()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
