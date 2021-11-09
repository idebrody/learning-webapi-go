package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

// sendParsing error back to the client
func sendParsingError(w http.ResponseWriter, qParam string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(`{"error": "Invalid number: %s"}`, qParam)))

}

type arithmetics func(num1 float64, num2 float64) float64

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func substract(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func division(num1 float64, num2 float64) float64 {
	//TODO: handle zero division
	return num1 / num2
}

func math(w http.ResponseWriter, r *http.Request, fn arithmetics) {

	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	//TODO: error handling
	num1, err := strconv.ParseFloat(query.Get("num1"), 64)
	if err != nil {
		sendParsingError(w, query.Get("num1"))
		return
	}

	num2, err := strconv.ParseFloat(query.Get("num2"), 64)
	if err != nil {
		sendParsingError(w, query.Get("num2"))
		return
	}

	w.WriteHeader(http.StatusOK)
	// trying to avoid trailing zeros
	w.Write([]byte(fmt.Sprintf(`{"result": %s}`, strconv.FormatFloat(fn(num1, num2), 'f', -1, 64))))
}

func notSupported(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"message": "not supported method"}`))
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) { math(w, r, add) }).Methods(http.MethodGet)
	api.HandleFunc("/substract", func(w http.ResponseWriter, r *http.Request) { math(w, r, substract) }).Methods(http.MethodGet)
	api.HandleFunc("/division", func(w http.ResponseWriter, r *http.Request) { math(w, r, division) }).Methods(http.MethodGet)
	// api.HandleFunc("/random", random).Methods(http.MethodGet)
	// api.HandleFunc("/metrics", metrics).Methods(http.MethodGet)
	// api.HandleFunc("/readiness", readiness).Methods(http.MethodGet)
	// api.HandleFunc("/liveness", liveness).Methods(http.MethodGet)
	api.HandleFunc("/", notSupported)
	log.Fatal(http.ListenAndServe(":8080", r))
}
