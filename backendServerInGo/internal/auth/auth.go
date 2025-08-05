package auth

import (
	"errors"   // To create and return custom error messages
	"net/http" // For working with HTTP headers
	"strings"  // For string splitting and checking parts of the Authorization header
)

// This function extracts the API key from the "Authorization" header in the HTTP request

// Example of expected header format:
// Authorization: ApiKey fsakfakfdkfdsadafdsa

// Parameters:
// - headers: http.Header → all the headers from the incoming HTTP request

// Returns:
// - string → the API key value if it’s found and valid
// - error  → error if the header is missing or incorrectly formatted
func GetAPIKey(headers http.Header) (string, error) {
	// Get the value of the "Authorization" header
	val := headers.Get("Authorization")

	// If no "Authorization" header is present, return an error
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	// Split the header value into two parts: expected to be "ApiKey actual_api_key"
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header") // Example: missing the key part
	}

	// Check if the first part is exactly "ApiKey"
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	// If everything is valid, return the actual API key (second part)
	return vals[1], nil
}
