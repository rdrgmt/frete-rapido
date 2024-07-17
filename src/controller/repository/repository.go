package repository

import (
	"fmt"
	"net/http"
)

// Welcome - a testpage to see if the server is running
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
