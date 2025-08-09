package main

import (
	"encoding/json" // For decoding JSON data from the request body
	"fmt"           // For formatting error messages
	"net/http"      // For handling HTTP requests and responses
	"time"          // For working with timestamps

	"github.com/Shariif-Mohsoo/ds-from-scratch-go/rssagg/internal/database" // Your internal database package
	"github.com/go-chi/chi"
	"github.com/google/uuid" // For generating unique user IDs
)

// This is a method of the `apiConfig` struct
// It handles feed creation when a request is made to the feed creation endpoint

// Parameters:
// - w: http.ResponseWriter → sends the response back to the client (like a browser or API tool)
// - r: *http.Request → represents the incoming HTTP request that contains user data
func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	// Define a struct to read JSON input from the request body
	// It expects a JSON object like: { "name": "Your Name" }
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
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
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),       // Generate a new unique ID
		CreatedAt: time.Now().UTC(), // Set current UTC time for when user was created
		UpdatedAt: time.Now().UTC(), // Set current UTC time for when user was updated
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		// If something went wrong while creating the user in DB, return an error
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	// If user creation was successful, return the user data as JSON with 200 OK status
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

// Parameters:
// - w: http.ResponseWriter → sends the response back to the client (like a browser or API tool)
// - r: *http.Request → represents the incoming HTTP request that contains user data
func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	// Call the CreateUser function from your database package to save the new user
	// Passing:
	// - a new unique ID for the user
	// - current time for CreatedAt and UpdatedAt
	// - the name received from the request
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		// If something went wrong while creating the user in DB, return an error
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follow: %v", err))
		return
	}

	// If user creation was successful, return the user data as JSON with 200 OK status
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}

// Parameters:
// - w: http.ResponseWriter → sends the response back to the client (like a browser or API tool)
// - r: *http.Request → represents the incoming HTTP request that contains user data
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		// If something went wrong while creating the user in DB, return an error
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		// If something went wrong while creating the user in DB, return an error
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
