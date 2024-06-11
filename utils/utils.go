package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// Parse the JSON payload from the request
func ParseJson(r *http.Request, payload any) error {
	// Check if the request body is empty
	if r.Body == nil {
		// If it is, return an error
		return fmt.Errorf("missing a request body")
	}
	// Decode the JSON payload
	return json.NewDecoder(r.Body).Decode(payload)
}

// Write the JSON response
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// Set the content type to JSON
	w.Header().Add("Content-Type", "application/json")
	// Set the status code
	w.WriteHeader(status)
	// Encode and return the JSON response
	return json.NewEncoder(w).Encode(v)
}

// Write the error response
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
