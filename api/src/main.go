package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/derchrischkya/async-api/api/src/endpoints"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//HTTP server for the api
	router := httprouter.New()

	// Technical Routes
	router.GET("/api/v1/ping", endpoints.Ping())

	router.GET("/api/v1/process/run", endpoints.Run())

	router.GET("/api/v1/process/state", endpoints.GetBackendProccessState())

	router.GET("/api/v1/internal/backendRun", endpoints.RunBackendProcess())

	port := 3000
	address := fmt.Sprintf(":%d", port)

	log.Println("Started: server is running on port", port)
	err := http.ListenAndServe(address, router)
	log.Fatal(err)
}
