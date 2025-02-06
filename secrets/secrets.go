package secrets

import (
	"os"
)

func GetAPIKey() string {
	apiKey := os.Getenv("TMDB_API_KEY") 
	if apiKey == "" {
		apiKey = "a276a946cd006baa84781efcf32d35e" 
	}
	return apiKey
}

func GetAccessToken() string {
	accessToken := os.Getenv("TMDB_ACCESS_TOKEN")
	if accessToken == "" {
		accessToken = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJhMjc2YTk0NmNkMDA2YmFhODQ3ODFlZmNmMzJkMzVlYyIsIm5iZiI6MTY5MDQ0NTE4NC45ODUsInN1YiI6IjY0YzIyNTgwMWNmZTNhMGViMjgyZDYxMiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.3CukNvquPzoW4ehXRbnNUpGu7Xo7YNC_HpYanxk9MUU" // Fallback for local testing (DO NOT COMMIT)
	}
	return accessToken
}
