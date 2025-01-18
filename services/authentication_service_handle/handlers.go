package authenticationservicehandle

import (
	"net/http"
)

// HomeHandler handles requests to the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}
