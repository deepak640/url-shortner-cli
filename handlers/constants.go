package handlers

import (
	"os"
	"github.com/joho/godotenv"
)

// DefaultServerURL is the production API endpoint.
const DefaultServerURL = "https://url-shortner-rosy-omega.vercel.app/"

// GetServerURL returns the server URL from environment variables or falls back to the default.
func GetServerURL() string {
	// Load .env if it exists (useful for local development or custom setups)
	godotenv.Load()

	server := os.Getenv("SERVER")
	if server == "" {
		return DefaultServerURL
	}

	// Ensure it ends with /
	if server[len(server)-1] != '/' {
		server += "/"
	}
	return server
}
