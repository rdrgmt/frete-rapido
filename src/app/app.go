package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter().StrictSlash(true)
)

// Init -
func Init() {
	// map the urls
	urlMaps()
	// start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
