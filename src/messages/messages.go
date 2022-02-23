package messages

import (
	"encoding/json"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	Response(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
