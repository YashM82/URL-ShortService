package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      string      `json:"error"`
}

func ResponseHandler(w http.ResponseWriter, statusCode int, message string, data interface{}, err error) {

	var apiResponse ApiResponse
	if err != nil {
		log.Println(message+": ", err)
		apiResponse = ApiResponse{StatusCode: statusCode, Message: message}

	} else {
		log.Println(message)
		apiResponse = ApiResponse{StatusCode: statusCode, Message: message, Data: data}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(apiResponse)
}
