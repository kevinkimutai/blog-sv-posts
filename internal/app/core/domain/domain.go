package domain

import (
	"errors"

	"github.com/kevinkimutai/metadata/internal/adapter/db/db"
)

type MovieResponse struct {
	StatusCode uint     `json:"status_code"`
	Message    string   `json:"message"`
	Data       db.Movie `json:"data"`
}

type MovieErrorResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
}

type MoviesResponse struct {
	StatusCode    uint       `json:"status_code"`
	Message       string     `json:"message"`
	Page          uint       `json:"page"`
	NumberOfPages uint       `json:"number_of_pages"`
	TotalMovies   uint       `json:"total_movies"`
	Data          []db.Movie `json:"data"`
}

type MovieParams struct{}

func NewMovieDomain(movie db.Movie) (db.Movie, error) {
	if movie.Title == "" || movie.Director == "" || movie.Description == "" {
		return movie, errors.New("missing title/director/description fields")
	}
	return movie, nil
}
