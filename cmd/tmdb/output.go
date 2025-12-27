package main

import (
	"fmt"
	"strings"

	"github.com/degerahmet/tmdb-cli/internal/tmdb"
)

func printMovies(movieType string, movies []tmdb.Movie, limit int) {
	label := map[string]string{
		"playing":  "Now Playing",
		"popular":  "Popular",
		"top":      "Top Rated",
		"upcoming": "Upcoming",
	}[movieType]

	if label == "" {
		label = "Movies"
	}

	if limit <= 0 {
		limit = 10
	}

	fmt.Printf("\n%s (showing up to %d)\n", label, limit)
	fmt.Println(strings.Repeat("-", 40))

	if len(movies) == 0 {
		fmt.Println("No movies found.")
		fmt.Println()
		return
	}

	if limit > len(movies) {
		limit = len(movies)
	}

	for i := 0; i < limit; i++ {
		m := movies[i]
		date := strings.TrimSpace(m.ReleaseDate)
		if date == "" {
			date = "N/A"
		}
		fmt.Printf("%2d) %s (%s) — Rating: %.1f\n", i+1, m.Title, date, m.VoteAverage)
	}
	fmt.Println()
}
