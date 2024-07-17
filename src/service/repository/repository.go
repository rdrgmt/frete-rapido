package service

import (
	"frete-rapido/src/domain/repository"
	"strconv"
	"strings"
)

// Check - validates the request
func Check(request repository.RequestQuote) (args []string) {
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
func Build(requestQuote repository.RequestQuote) (requestAPI repository.RequestAPI) {
	// shipper
	requestAPI.Shipper.RegisteredNumber = "25438296000158"
	requestAPI.Shipper.Token = "1d52a9b6b78cf07b08586152459a5c90"
	requestAPI.Shipper.PlatformCode = "5AKVkHqCn"

	// recipient
	requestAPI.Recipient.Type = 0
	requestAPI.Recipient.Country = "BRA" // fixed for now
	requestAPI.Recipient.Zipcode, _ = strconv.Atoi(requestQuote.Recipient.Address.Zipcode)

	// dispatchers
	var dispatcher repository.Dispatcher
	dispatcher.RegisteredNumber = "25438296000158"
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
