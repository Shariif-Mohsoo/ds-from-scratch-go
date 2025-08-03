package main

import "net/http" // Used to handle HTTP requests and responses

// This function handles a "readiness" check — usually used in health checks by servers or load balancers
// It simply responds with a success message (200 OK)

// Parameters:
// - w: http.ResponseWriter → used to send the response back to the client
// - r: *http.Request → represents the incoming request from the client (not used here)
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// This sends a 200 OK status with an empty JSON object: {}
	// struct{}{} → defines an empty struct with no fields (used to return an empty JSON object)
	respondWithJSON(w, 200, struct{}{})
}
