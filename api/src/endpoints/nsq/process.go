package nsq

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func StartAsyncProcess() httprouter.Handle {
	return func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		log.Println("Started: Async process")
		time.Sleep(10000000000)
		writer.WriteHeader(http.StatusOK)
	}
}
