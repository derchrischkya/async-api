package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/derchrischkya/async-api/api/src/endpoints"
	"github.com/derchrischkya/async-api/api/src/endpoints/nsq"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//HTTP server for the api
	router := httprouter.New()

	// Technical Routes
	router.GET("/api/v1/ping", endpoints.Ping())

	router.GET("/api/v1/nsq/public/run", nsq.Run())

	router.GET("/api/v1/nsq/public/state", nsq.GetAsyncProcessState())

	router.GET("/api/v1/nsq/internal/backendRun", nsq.StartAsyncProcess())

	port := 3000
	address := fmt.Sprintf(":%d", port)

	log.Println("Started: server is running on port", port)
	err := http.ListenAndServe(address, router)
	log.Fatal(err)
}
