package controllers

import (
	"URL-ShortService/constants"
	"URL-ShortService/db"
	"URL-ShortService/util"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-playground/validator"
)

type Input struct {
	LongURL string `json:"longURL" validate:"required"`
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		util.ResponseHandler(w, constants.StatusCodes.INTERNAL_SERVER_ERROR, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		util.ResponseHandler(w, constants.StatusCodes.BAD_REQUEST, "Invalid Input", nil, err)
		return
	}

	// Validate the format of the long URL
	parsedURL, err := url.ParseRequestURI(input.LongURL)
	if err != nil || parsedURL.Scheme == "" {
		util.ResponseHandler(w, constants.StatusCodes.BAD_REQUEST, "Invalid URL format", nil, err)
		return
	}

	shortURL := generateShortURL()
	if shortURL == "" {
		util.ResponseHandler(w, constants.StatusCodes.INTERNAL_SERVER_ERROR, "Failed to generate short URL", nil, nil)
		return
	}
	bindedURL := "http://localhost:8080/" + shortURL
	// Save the short and long URL in the database
	err = saveURLMapping(input.LongURL, bindedURL, shortURL)
	if err != nil {
		util.ResponseHandler(w, constants.StatusCodes.INTERNAL_SERVER_ERROR, "Could not save URL mapping", nil, err)
		return
	}

	// Return the short URL to the client
	response := map[string]string{"shortURL": fmt.Sprintf(bindedURL)}
	util.ResponseHandler(w, constants.StatusCodes.CREATED, "URL Shortened Successfully", response, nil)

}

func saveURLMapping(longURL string, shortURL string, shortURLcode string) error {
	query := "INSERT INTO urls (long_url, short_url,short_code) VALUES (?, ?,?)"
	stmt, err := db.MySqlSession.Prepare(query)
	if err != nil {
		log.Println("Error preparing SQL statement:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(longURL, shortURL, shortURLcode)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		return err
	}
	return nil
}

// generateShortURL generates a random 6-character string for the short URL
func generateShortURL() string {
	const urlLength = 6
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, urlLength)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Error generating random bytes:", err)
		return " error"
	}

	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b)
}
