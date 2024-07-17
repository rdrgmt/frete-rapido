package app

import (
	"frete-rapido/src/controller/repository"
)

// urlMaps - maps the urls to the respective functions
func urlMaps() {
	router.HandleFunc("/", repository.Welcome)
	router.HandleFunc("/quote", repository.Quote).Methods("POST")
}
