package ports

import (
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
)

type DBPort interface {
	//Movie Interfaces
	GetMovieById(movieID int64) (*domain.Movie, error)
	CreateMovie(movie domain.Movie) (domain.Movie, error)
	//Ratings Interfaces
	CreateRating(rating domain.Rating) (domain.Rating, error)
}
