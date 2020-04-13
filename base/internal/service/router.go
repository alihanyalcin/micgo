package {servicename}

import (
	"github.com/gorilla/mux"
	"net/http"
)

func LoadRestRoutes() *mux.Router {
	r := mux.NewRouter()

	b := r.PathPrefix("/api/v1").Subrouter()

	// Ping Resource
	b.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	return r
}

// pingHandler implements a controller to execute a ping request.
func pingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("pong"))
}
