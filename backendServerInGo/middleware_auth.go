package main

import (
	"fmt"      // For formatting error messages
	"net/http" // For handling HTTP requests and responses

	"github.com/Shariif-Mohsoo/ds-from-scratch-go/rssagg/internal/auth"     // For authentication functions like extracting API key
	"github.com/Shariif-Mohsoo/ds-from-scratch-go/rssagg/internal/database" // For database functions like getting user by API key
)

// Define a custom type for handlers that require authentication
// authedHandler is a function that takes:
// - http.ResponseWriter → to send response back
// - *http.Request → the incoming request
// - database.User → the user that was authenticated successfully
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// This is a method of apiConfig called middlewareAuth
// It wraps around routes (handlers) that need user authentication using API key

// Parameters:
// - handler: authedHandler → the actual handler function to run *after* user is authenticated

// Returns:
// - http.HandlerFunc → a standard Go handler that can be used in routing (like http.HandleFunc)
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	// Return a new handler function that includes authentication logic
	return func(w http.ResponseWriter, r *http.Request) {
		// Try to extract API key from the request header
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			// If API key is missing or invalid, return 403 Forbidden error
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		// Use the API key to find the corresponding user in the database
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			// If user is not found or error occurs, return 400 Bad Request
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		// If everything is successful, call the original handler function
		// and pass the authenticated user to it
		handler(w, r, user)
	}
}
