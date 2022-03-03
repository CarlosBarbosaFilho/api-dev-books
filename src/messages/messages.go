package messages

import (
	"encoding/json"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, error error) {
	Response(w, statusCode, struct {
		Err string `json:"err"`
	}{
		Err: error.Error(),
	})
}

func ResponseGeneric(err error, message interface{}) {
	message = "calling error"
	if err != nil {
		log.Fatal(err)
	}
}

func GenericError(err error, message string) {
	ResponseGeneric(err, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
