package pkg

import (
	"encoding/json"
	"net/http"
	"{project}/internal/pkg/logger"
)

func Encode(i interface{}, w http.ResponseWriter, LoggingClient logger.LoggingClient) {
	w.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err := enc.Encode(i)
	if err != nil {
		LoggingClient.Error("Error encoding the data: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
