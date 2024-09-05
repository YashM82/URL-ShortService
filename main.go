package main

import (
	"URL-ShortService/db"
	"URL-ShortService/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func makeConnections() {
	// Initialize the database connection
	db.ConnectMySqlDb()
}
func main() {

	makeConnections()

	// Initialize the Gorilla Mux router
	r := mux.NewRouter()

	// Register routes from different files
	routes.RegisterURLRoutes(r)	
	

	// Start the server

	log.Println("Server is running on port: 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
