package nsq

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetAsyncProcessState() httprouter.Handle {
	return func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotImplemented)
		writer.Write([]byte(`{"message": "Not implemented"}`))
	}
}
