package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/degerahmet/tmdb-cli/internal/cli"
	"github.com/degerahmet/tmdb-cli/internal/config"
	"github.com/degerahmet/tmdb-cli/internal/tmdb"
)

func main() {
	config.Load()

	flag.Usage = cli.PrintUsage

	typeFlag := flag.String("type", "popular", "movie type")
	limitFlag := flag.Int("limit", 10, "limit")
	queryFlag := flag.String("query", "", "search query")
	flag.Parse()

	var endpoint string
	var err error

	if *queryFlag != "" {
		endpoint = "/search/movie?query=" + url.QueryEscape(*queryFlag)
	} else {
		endpoint, err = getEndpoint(*typeFlag)
		if err != nil {
			fmt.Println("Error:", err)
			cli.PrintUsage()
			os.Exit(1)
		}
	}

	apiKey, err := config.TMDBApiKey()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	movies, err := tmdb.FetchMovies(apiKey, endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var moviesLabel string
	if *queryFlag != "" {
		moviesLabel = fmt.Sprintf("Search Results for %q", *queryFlag)
	} else {
		moviesLabel = strings.ToLower(*typeFlag)
	}

	printMovies(moviesLabel, movies, *limitFlag)
}
