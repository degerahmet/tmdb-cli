package main

import (
	"flag"
	"fmt"
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
	flag.Parse()

	endpoint, err := getEndpoint(*typeFlag)
	if err != nil {
		fmt.Println("Error:", err)
		cli.PrintUsage()
		os.Exit(1)
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

	printMovies(strings.ToLower(*typeFlag), movies, *limitFlag)
}
