package tmdb

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
}
type MovieResponse struct {
	Page    int     `json:"page"`
	Results []Movie `json:"results"`
}
