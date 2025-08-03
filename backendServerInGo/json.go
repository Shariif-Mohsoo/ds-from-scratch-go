package main

import (
	"encoding/json" // For converting Go data to JSON format
	"log"           // For printing log messages (errors, info, etc.)
	"net/http"      // For creating HTTP servers and handling requests/responses
)

// This function sends an error response in JSON format
// Parameters:
// - w: http.ResponseWriter → used to send the response back to the client (like browser or API user)
// - code: int → HTTP status code (like 400 for bad request, 500 for server error, etc.)
// - msg: string → the error message that we want to send to the client
func respondWithError(w http.ResponseWriter, code int, msg string) {
	// If the status code is 500 or above (which means server error), print the error in logs
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	// Creating a structure (struct) to hold the error message in JSON format
	// The resulting JSON will look like: {"error": "your error message"}
	type errResponse struct {
		Error string `json:"error"` // This makes the field name appear as "error" in the JSON output
	}

	// Calling another function to send the error response in JSON format
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

// This function sends any data as a JSON response
// Parameters:
// - w: http.ResponseWriter → used to send the JSON response to the client
// - code: int → the HTTP status code (like 200 for success, 404 for not found)
// - payload: interface{} → the actual data you want to send (can be any type like string, struct, map, etc.)
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Convert the payload data into JSON format
	data, err := json.Marshal(payload)
	if err != nil {
		// If there's an error in conversion, log it and return 500 server error
		log.Printf("\nFailed to marshal JSON response: %v", payload)
		w.WriteHeader(500) // Internal server error
		return
	}

	// Tell the browser or client that we are sending JSON data
	w.Header().Add("Content-Type", "application/json")

	// Set the status code for the response (like 200, 400, etc.)
	w.WriteHeader(code)

	// Write (send) the final JSON data to the client
	w.Write(data)
}
