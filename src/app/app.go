package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	mongodb "frete-rapido/src/db/repository"
)

var (
	router = mux.NewRouter().StrictSlash(true)
)

// Init -
func Init() {
	// map the urls
	urlMaps()

	// create the database
	mongodb.CreateDB()

	// start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
