package models

type Movie struct {
	Title       string  `json:"title"`
	Rating      float32 `json:"vote_average"`
	ReleaseDate string  `json:"release_date"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}
