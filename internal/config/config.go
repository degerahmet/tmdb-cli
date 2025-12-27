package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Load() {
	// Load .env file if it exists
	_ = godotenv.Load()
}

func TMDBApiKey() (string, error) {
	key := strings.TrimSpace(os.Getenv("TMDB_API_KEY"))
	if key == "" {
		return "", fmt.Errorf("TMDB_API_KEY is not set")
	}
	return key, nil
}
