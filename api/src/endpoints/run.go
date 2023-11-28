package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type Response struct {
	ID      string  `json:"id"`
	Created bool    `json:"created"`
	Message Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Href string `json:"_href"`
}

func Run() httprouter.Handle {
	return func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		uuid := uuid.New()
		data := []byte(fmt.Sprintf("{\"session_id\": \"%s\"}", uuid.String()))
		log.Println("Started: Login to server with session ID", uuid.String())

		// Keep session for 10 minutes and then logout
		// Logout is done by another endpoint, event message system handles the logout
		resp, _ := http.Post("http://127.0.0.1:4151/pub?topic=async-api&channel=backendDemoProcessor", "text/plain", bytes.NewBuffer(data))
		bodyString := ""
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString = string(bodyBytes)
			log.Println(bodyString)
		}

		// For various message-streaming systems, we can use diffrent response body to send back a message
		// Some of them may offer a backtrace url to check the status of the process, others just fire and forget
		message := Message{
			Text: "Process started",
			Href: ""}

		if bodyString != "OK" {
			message = Message{
				Text: "Process started",
				Href: "mylink"}
		}

		responseBody := Response{
			ID:      uuid.String(),
			Created: true,
			Message: message}

		// Marshal the response body to JSON
		responseBodyBytes, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(writer, "Failed to insert book into the database", http.StatusInternalServerError)
			log.Printf("Failed to marshal response: %v", err)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusAccepted)
		writer.Write(responseBodyBytes)
	}
}
