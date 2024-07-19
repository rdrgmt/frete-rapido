package app

import (
	"log"
	"net/http"

	controller "frete-rapido/src/controller"
	mongodb "frete-rapido/src/db"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter().StrictSlash(true)
)

// Init -
func Init() {
	// create the database
	mongodb.CreateDB()

	// map the urls
	router.HandleFunc("/", controller.Welcome)
	router.HandleFunc("/quote", controller.Quote).Methods("POST")
	// start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
