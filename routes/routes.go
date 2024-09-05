package routes

import (
	abc "URL-ShortService/controllers"

	"github.com/gorilla/mux"
)

func RegisterURLRoutes(r *mux.Router) {
	r.HandleFunc("/shortService", abc.CreateShortURL).Methods("POST")
	// Catch-all route for redirecting short URLs
	r.PathPrefix("/{shortCode}").HandlerFunc(abc.RedirectToLongURL).Methods("GET")

}
