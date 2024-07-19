package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"frete-rapido/src/config"
	mongodb "frete-rapido/src/db"
	domain "frete-rapido/src/domain"
	"net/http"
	"strconv"
	"strings"
)

// Check - validates the request
func Check(request domain.RequestQuote) (args []string) {
	args = make([]string, 0)

	// contains zipcode?
	if strings.EqualFold(request.Recipient.Address.Zipcode, "") {
		args = append(args, "Zipcode is required")
	}

	// contains volumes?
	if len(request.Volumes) == 0 {
		args = append(args, "Volumes is required")
	}

	// contains specific variables?
	for i, volume := range request.Volumes {
		if strings.EqualFold(volume.Category, "") {
			args = append(args, "Category is required for Volume "+strconv.Itoa(i))
		}
		if volume.Amount == 0 {
			args = append(args, "Amount is required for Volume "+strconv.Itoa(i))
		}
		if volume.UnitaryWeight == 0 {
			args = append(args, "UnitaryWeight is required for Volume "+strconv.Itoa(i))
		}
		if volume.Price == 0 {
			args = append(args, "Price is required for Volume "+strconv.Itoa(i))
		}
		if strings.EqualFold(volume.Sku, "") {
			args = append(args, "Sku is required for Volume "+strconv.Itoa(i))
		}
		if volume.Height == 0 {
			args = append(args, "Height is required for Volume "+strconv.Itoa(i))
		}
		if volume.Width == 0 {
			args = append(args, "Width is required for Volume "+strconv.Itoa(i))
		}
		if volume.Length == 0 {
			args = append(args, "Length is required for Volume "+strconv.Itoa(i))
		}
	}

	return args
}

// Build - builds the request to the API
func Build(requestQuote domain.RequestQuote) (requestAPI domain.RequestAPI) {
	// shipper
	requestAPI.Shipper.RegisteredNumber = config.RegisteredNumber
	requestAPI.Shipper.Token = config.Token
	requestAPI.Shipper.PlatformCode = config.PlatformCode

	// recipient
	requestAPI.Recipient.Type = 0        // fixed for now
	requestAPI.Recipient.Country = "BRA" // fixed for now
	requestAPI.Recipient.Zipcode, _ = strconv.Atoi(requestQuote.Recipient.Address.Zipcode)

	// dispatchers
	var dispatcher domain.Dispatcher
	dispatcher.RegisteredNumber = config.RegisteredNumber
	dispatcher.Zipcode = requestAPI.Recipient.Zipcode
	for _, volume := range requestQuote.Volumes {
		volume.UnitaryPrice = volume.Price / volume.Amount
		dispatcher.Volumes = append(dispatcher.Volumes, volume)
	}
	requestAPI.Dispatchers = append(requestAPI.Dispatchers, dispatcher)

	// simulation type
	requestAPI.SimulationType = append(requestAPI.SimulationType, 0) // fixed for now

	// returns
	requestAPI.Returns.Composition = false  // fixed for now
	requestAPI.Returns.Volumes = false      // fixed for now
	requestAPI.Returns.AppliedRules = false // fixed for now

	return requestAPI
}

// Simulate - sends the request to the API
func Simulate(requestAPI domain.RequestAPI) (responseAPI domain.ResponseAPI, err error) {
	path := config.Path
	method := http.MethodPost

	// build the request
	payload, err := json.Marshal(requestAPI)
	if err != nil {
		return responseAPI, err
	}

	client := &http.Client{}
	request, err := http.NewRequest(method, path, bytes.NewReader(payload))
	if err != nil {
		return responseAPI, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	// send the request
	response, err := client.Do(request)
	if err != nil {
		return responseAPI, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return responseAPI, errors.New("API returned a non-200 status code")
	}

	// decode the response
	err = json.NewDecoder(response.Body).Decode(&responseAPI)
	if err != nil {
		return responseAPI, errors.New("Failed to decode the response: " + err.Error())
	}

	return responseAPI, err
}

// Format - formats the response from the API to the desired format
func Format(responseAPI domain.ResponseAPI) (responseQuote domain.ResponseQuote) {
	// check if there are dispatchers
	if len(responseAPI.Dispatchers) == 0 {
		return responseQuote
	}

	// format the response
	for _, offer := range responseAPI.Dispatchers[0].Offers {
		responseQuote.Carrier = append(responseQuote.Carrier, domain.CarrierQuote{
			Name:     offer.Carrier.Name,
			Service:  offer.Modal,
			Deadline: offer.CarrierOriginalDeliveryTime.Days,
			Price:    offer.FinalPrice,
		})
	}

	return responseQuote
}

// Prepare - prepares the quotes to be posted to the api
func Prepare(quotes []mongodb.QuoteBD) (responseMetrics domain.ResponseMetrics, err error) {
	// create map to store metrics
	mapMetrics := domain.Metric{
		ResultsPerCarrier:    make(map[string]int),
		TotalPricePerCarrier: make(map[string]float64),
		AvgPricePerCarrier:   make(map[string]float64),
		CheapestFreight:      make(map[string]float64),
		PriciestFreight:      make(map[string]float64),
	}

	// iterate over the quotes
	for _, quote := range quotes {
		for _, carrier := range quote.Carriers {
			// total results
			mapMetrics.ResultsPerCarrier[carrier.Name]++

			// total price
			mapMetrics.TotalPricePerCarrier[carrier.Name] = (mapMetrics.TotalPricePerCarrier[carrier.Name] + carrier.Price)
			mapMetrics.TotalPricePerCarrier[carrier.Name] = float64(int(mapMetrics.TotalPricePerCarrier[carrier.Name]*100)) / 100

			// average price
			mapMetrics.AvgPricePerCarrier[carrier.Name] = mapMetrics.TotalPricePerCarrier[carrier.Name] / float64(mapMetrics.ResultsPerCarrier[carrier.Name])
			mapMetrics.AvgPricePerCarrier[carrier.Name] = float64(int(mapMetrics.AvgPricePerCarrier[carrier.Name]*100)) / 100

			// cheapest freight
			if mapMetrics.CheapestFreight[carrier.Name] == 0 || carrier.Price < mapMetrics.CheapestFreight[carrier.Name] {
				mapMetrics.CheapestFreight[carrier.Name] = carrier.Price
			}

			// priciest freight
			if mapMetrics.PriciestFreight[carrier.Name] == 0 || carrier.Price > mapMetrics.PriciestFreight[carrier.Name] {
				mapMetrics.PriciestFreight[carrier.Name] = carrier.Price
			}
		}
	}

	// append the metrics to the response
	responseMetrics.Metrics = append(responseMetrics.Metrics, mapMetrics)

	return responseMetrics, err
}
