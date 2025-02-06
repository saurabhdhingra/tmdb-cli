package cmd

import (
	"fmt"
	"tmdb-app/models"
)

func DisplayMovies(movies []models.Movie) {
	if len(movies) == 0 {
		fmt.Println("No movies found.")
		return
	}

	fmt.Println("Movies List:")
	for _, movie := range movies {
		fmt.Printf("Title: %s | Rating: %f | Release Date: %s\n", movie.Title, movie.Rating, movie.ReleaseDate)
	}
}
