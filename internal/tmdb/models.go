package tmdb

type Movie struct {
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
}

type MovieResponse struct {
	Page    int     `json:"page"`
	Results []Movie `json:"results"`
}
