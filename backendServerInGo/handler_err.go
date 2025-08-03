package main

import "net/http" // For handling HTTP requests and responses

// This function is a request handler that sends an error response when called
// It can be used in an HTTP route like: http.HandleFunc("/somepath", handlerError)

// Parameters:
// - w: http.ResponseWriter → used to send the response back to the client (like browser or API user)
// - r: *http.Request → represents the incoming HTTP request from the client (not used in this example, but required)
func handlerError(w http.ResponseWriter, r *http.Request) {
	// Call the custom error response function and send a 400 (Bad Request) error
	// "Something went wrong" is the error message sent back to the client in JSON format
	respondWithError(w, 400, "Something went wrong")
}
