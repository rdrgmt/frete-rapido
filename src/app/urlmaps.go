package app

import (
	"frete-rapido/src/controller/repository"
	"log"
	"net/http"
)

// urlMaps - maps the urls to the respective functions
func urlMaps() {
	router.HandleFunc("/", repository.Welcome)
	log.Fatal(http.ListenAndServe(":8080", router))
}
