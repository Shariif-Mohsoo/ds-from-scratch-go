package main

import (
	"context"
	"encoding/json" // For decoding JSON data from the request body
	"fmt"           // For formatting error messages
	"net/http"      // For handling HTTP requests and responses
	"time"          // For working with timestamps

	"github.com/Shariif-Mohsoo/ds-from-scratch-go/rssagg/internal/database" // Your internal database package
	"github.com/google/uuid"                                                // For generating unique user IDs
)

// This is a method of the `apiConfig` struct
// It handles user creation when a request is made to the user creation endpoint

// Parameters:
// - w: http.ResponseWriter → sends the response back to the client (like a browser or API tool)
// - r: *http.Request → represents the incoming HTTP request that contains user data
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Define a struct to read JSON input from the request body
	// It expects a JSON object like: { "name": "Your Name" }
	type parameters struct {
		Name string `json:"name"` // This tag tells Go to map the "name" key from JSON to the Name field
	}

	// Create a new JSON decoder to read the request body
	decoder := json.NewDecoder(r.Body)

	// Create an empty `parameters` struct to hold the decoded data
	params := parameters{}

	// Try to decode the JSON body into the `params` struct
	err := decoder.Decode(&params)
	if err != nil {
		// If decoding fails, return a 400 Bad Request error with details
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}

	// Call the CreateUser function from your database package to save the new user
	// Passing:
	// - a new unique ID for the user
	// - current time for CreatedAt and UpdatedAt
	// - the name received from the request
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),       // Generate a new unique ID
		CreatedAt: time.Now().UTC(), // Set current UTC time for when user was created
		UpdatedAt: time.Now().UTC(), // Set current UTC time for when user was updated
		Name:      params.Name,      // Set user's name from the request
	})

	if err != nil {
		// If something went wrong while creating the user in DB, return an error
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	// If user creation was successful, return the user data as JSON with 200 OK status
	respondWithJSON(w, 201, databaseUserToUser(user))
}

// This function handles a request to get user information (usually used in authenticated routes)
// It returns the user data in JSON format to the client

// Parameters:
// - apiCfg: *apiConfig → your app's configuration, giving access to database and other settings
// - w: http.ResponseWriter → used to send the response back to the client (like a browser or API client)
// - r: *http.Request → the incoming HTTP request (not used here, but required)
// - user: database.User → the user who has already been authenticated (passed from middleware)
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// Convert the database user format to a clean user format (for JSON)
	// Then send it as a JSON response with status code 200 (OK)
	respondWithJSON(w, 200, databaseUserToUser(user))
}

// Parameters:
// - apiCfg: *apiConfig → your app's configuration, giving access to database and other settings
// - w: http.ResponseWriter → used to send the response back to the client (like a browser or API client)
// - r: *http.Request → the incoming HTTP request (not used here, but required)
// - user: database.User → the user who has already been authenticated (passed from middleware)
func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}
	respondWithJSON(w, 200, databasePostsToPosts(posts))

}
