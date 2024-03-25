package domain

import (
	"errors"
	"time"
)

type DataResponse struct {
	StatusCode uint        `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
}

type Response struct {
	StatusCode    uint          `json:"status_code"`
	Message       string        `json:"message"`
	Page          uint          `json:"page"`
	NumberOfPages uint          `json:"number_of_pages"`
	TotalMovies   uint          `json:"total_movies"`
	Data          []interface{} `json:"data"`
}

type MovieParams struct{}

type Movie struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Director      string    `json:"director"`
	CreatedAt     time.Time `json:"created_at"`
	AverageRating float64   `json:"average_rating"`
}

type Rating struct {
	ID        int64     `json:"id"`
	MovieID   int64     `json:"movie_id"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMovieDomain(movie Movie) (Movie, error) {
	if movie.Title == "" {
		return movie, errors.New("missing title field")
	}
	if movie.Description == "" {
		return movie, errors.New("missing description field")
	}
	if movie.Director == "" {
		return movie, errors.New("missing director field")
	}

	return movie, nil
}

func NewRatingsDomain(rating Rating) (Rating, error) {
	if rating.MovieID == 0 {
		return rating, errors.New("missing movie_id field")
	}
	if rating.Rating == 0 {
		return rating, errors.New("missing rating field")
	}

	return rating, nil
}
