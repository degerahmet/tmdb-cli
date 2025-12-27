package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.themoviedb.org/3"

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

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(u.String())
	if err != nil {
		return nil, err
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
