package repository

import (
	"encoding/json"
	"fmt"
	domain "frete-rapido/src/domain/repository"
	service "frete-rapido/src/service/repository"
	"net/http"
)

// Welcome - a testpage to see if the server is running
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

// Quote -
func Quote(w http.ResponseWriter, r *http.Request) {
	request := domain.RequestQuote{}

	// decode the request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request: Body is not a Valid JSON"})
		return
	}

	// validate the request
	invalidArgs := service.Check(request)
	if len(invalidArgs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string][]string{"error": invalidArgs})
		return
	}

	// build the request
	requestAPI := service.Build(request)

	// simulate the request
	responseAPI, err := service.Simulate(requestAPI)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// format the response
	responseQuote := service.Format(responseAPI)

	// all set! return the request
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseQuote)

	// TODO - save the request to the database

}
