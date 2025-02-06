package cmd

import (
	"flag"
	"fmt"
)

func Execute() {
	category := flag.String("type", "", "Specify movie category: playing, popular, top, upcoming")
	flag.Parse()

	if *category == "" {
		fmt.Println("Usage: tmdb-app --type \"category\" (playing, popular, top, upcoming)")
		return
	}

	movies, err := FetchMovies(*category)
	if err != nil {
		fmt.Println("Errors:", err)
		return
	}

	DisplayMovies(movies)
}
