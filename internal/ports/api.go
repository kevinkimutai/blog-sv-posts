package ports

import (
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
)

type APIPort interface {
	//Movie Interfaces
	GetMovie(movieID int64) (*domain.Movie, error)
	CreateNewMovie(movie domain.Movie) (domain.Movie, error)
	//Ratings Interfaces
	CreateNewRating(rating domain.Rating) (domain.Rating, error)
}
