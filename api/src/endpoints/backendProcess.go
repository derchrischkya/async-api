package endpoints

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func GetBackendProccessState() httprouter.Handle {
	return func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		uuid := uuid.New()
		data := []byte(fmt.Sprintf("{\"session_id\": \"%s\"}", uuid.String()))
		log.Println("Started: Login to server with session ID", uuid.String())

		// Keep session for 10 minutes and then logout
		// Logout is done by another endpoint, event message system handles the logout
		err, _ := http.Post("http://127.0.0.1:4151/pub?topic=async-api&channel=logout", "text/plain", bytes.NewBuffer(data))

		if err != nil {
			log.Println("Error: Async message could not be sent to logout topic")
			log.Println(err)
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func RunBackendProcess() httprouter.Handle {
	return func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		time.Sleep(10000000000)
		writer.WriteHeader(http.StatusOK)
	}
}
