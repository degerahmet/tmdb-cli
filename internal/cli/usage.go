package cli

import "fmt"

func PrintUsage() {
	fmt.Println("tmdb - A CLI tool to fetch movie data from TMDB")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  tmdb --type <type> [--limit <n>]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --type   playing | popular | top | upcoming")
	fmt.Println("  --limit  number of movies to display")
	fmt.Println("  -h, --help  show help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  tmdb --type popular --limit 5")
}
