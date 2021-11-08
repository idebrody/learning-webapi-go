package main

import (
    "log"
    "net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
}

func notSupported(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotImplemented)
    w.Write([]byte(`{"message": "not supported method"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
    r.HandleFunc("/", notSupported)
    log.Fatal(http.ListenAndServe(":8080", r))
}