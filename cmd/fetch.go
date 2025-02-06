package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"tmdb-app/models"
	"tmdb-app/secrets"
)

var endpoints = map[string]string{
	"playing":  "https://api.themoviedb.org/3/movie/now_playing",
	"popular":  "https://api.themoviedb.org/3/movie/popular",
	"top":      "https://api.themoviedb.org/3/movie/top_rated",
	"upcoming": "https://api.themoviedb.org/3/movie/upcoming",
}

func FetchMovies(category string) ([]models.Movie, error) {
	accessToken := secrets.GetAccessToken()
	if accessToken == "" {
		return nil, errors.New("access token is missing. Set TMDB_ACCESS_TOKEN environment variable")
	}

	url, exists := endpoints[category]
	if !exists {
		return nil, errors.New("invalid category. Use: playing, popular, top, upcoming")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var movieResponse models.MovieResponse
	if err := json.Unmarshal(body, &movieResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return movieResponse.Results, nil
}
