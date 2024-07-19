package repository

import (
	"encoding/json"
	"fmt"
	mongodb "frete-rapido/src/db"
	domain "frete-rapido/src/domain"
	service "frete-rapido/src/service"
	"net/http"
	"strconv"
	"strings"
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

	// save the request to the database
	err = mongodb.SaveQuoteDB(responseQuote)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// all set! return the request
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseQuote)
}

// Metrics - retrieve the last quotes from the database
func Metrics(w http.ResponseWriter, r *http.Request) {
	var lastQuotes int64
	var err error

	// check if the querystring is set
	queryString := r.URL.Query().Get("last_quotes")
	if !strings.EqualFold(queryString, "") {
		lastQuotes, err = strconv.ParseInt(queryString, 10, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request: last_quotes is not a valid number"})
			return
		}
	}

	// retrieve the quotes from the database
	quotes, err := mongodb.RetrieveQuotesDB(lastQuotes)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// prepare the response
	responseMetrics, err := service.Prepare(quotes)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// all set! return the request
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseMetrics)
}
