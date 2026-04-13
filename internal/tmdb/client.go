package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.themoviedb.org/3"

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func FetchMovies(apiKey, endpoint string) ([]Movie, error) {
	u, err := url.Parse(baseURL + endpoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("api_key", apiKey)
	q.Set("language", "en-US")
	q.Set("page", "1")
	u.RawQuery = q.Encode()

	var resp *http.Response
	var lastErr error

	// Retry loop for transient network/TCP errors (up to 3 attempts)
	for i := 0; i < 3; i++ {
		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return nil, err
		}

		// Set User-Agent and Accept headers
		req.Header.Set("User-Agent", "tmdb-cli/1.0")
		req.Header.Set("Accept", "application/json")

		resp, lastErr = client.Do(req)
		if lastErr == nil {
			break
		}

		// Wait before retrying (exponential backoff)
		time.Sleep(time.Duration(i+1) * 500 * time.Millisecond)
	}

	if lastErr != nil {
		return nil, fmt.Errorf("network error after retries: %w", lastErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tmdb api error: %s", resp.Status)
	}

	var data MovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Results, nil
}
