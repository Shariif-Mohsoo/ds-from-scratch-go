package main

import (
	"time" // For time-related data types like time.Time

	"github.com/Shariif-Mohsoo/ds-from-scratch-go/rssagg/internal/database" // Importing your internal database package
	"github.com/google/uuid"                                                // For using UUIDs (unique user IDs)
)

// This struct defines what the user data will look like when returned as JSON
// It is used to send clean and structured user information to the client
type User struct {
	ID        uuid.UUID `json:"id"`         // Unique user ID (example: "4f4a1e7a-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
	CreatedAt time.Time `json:"created_at"` // Time when the user was created
	UpdatedAt time.Time `json:"updated_at"` // Time when the user was last updated
	Name      string    `json:"name"`       // The user's name
	APIKey    string    `json:"api_key"`    //The api key for auth
}

// This function converts a user from the database format to the format we want to send to the client
// Parameter:
// - dbUser: database.User → the user data returned from the database
// Returns:
// - User: the cleaned-up user data ready to be sent as a JSON response
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,        // Copy the user's ID from the database result
		CreatedAt: dbUser.CreatedAt, // Copy the user's created time
		UpdatedAt: dbUser.UpdatedAt, // Copy the user's updated time
		Name:      dbUser.Name,      // Copy the user's name
		APIKey:    dbUser.ApiKey,
	}
}
