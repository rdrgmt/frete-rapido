package repository

import (
	"encoding/json"
	"fmt"
	"frete-rapido/src/domain/repository"
	"net/http"
	"strings"
)

// Welcome - a testpage to see if the server is running
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

// Quote -
func Quote(w http.ResponseWriter, r *http.Request) {
	request := repository.RequestQuote{}

	// decode the request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request: Body is not a Valid JSON"})
		return
	}

	// validate the request
	// contains zipcode?
	if strings.EqualFold(request.Recipient.Address.Zipcode, "") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request: Zipcode is required"})
		return
	}

	// contains volumes?
	if len(request.Volumes) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request: Volumes is required"})
		return
	}

	// all set! return the request
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(request)
}
