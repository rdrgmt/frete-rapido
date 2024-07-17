package app

import (
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter().StrictSlash(true)
)

// StartApp -
func StartApp() {
	urlMaps()

}
