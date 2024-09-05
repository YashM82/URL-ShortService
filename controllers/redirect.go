package controllers

import (
	"URL-ShortService/constants"
	"URL-ShortService/db"
	"URL-ShortService/util"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RedirectToLongURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	log.Println("Received request for shortCode:", shortCode)

	// Query the database for the long URL
	longURL, err := getLongURLFromDB(shortCode)
	if err != nil {
		log.Println("Error retrieving long URL for shortCode:", shortCode)
		util.ResponseHandler(w, constants.StatusCodes.NOT_FOUND, "URL not found", nil, err)
		return
	}
	log.Println("Redirecting to:", longURL)
	// Redirect to the long URL
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

func getLongURLFromDB(shortCode string) (string, error) {
	var longURL string
	query := "SELECT long_url FROM urls WHERE short_code = ?"
	err := db.MySqlSession.QueryRow(query, shortCode).Scan(&longURL)
	if err != nil {
		log.Println("Error fetching long URL from DB:", err)
		return "", err
	}
	return longURL, nil
}
