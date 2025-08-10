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

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishAt   time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishAt:   dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost((dbPost)))
	}
	return posts
}
