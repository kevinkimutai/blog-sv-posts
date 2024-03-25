package application

import (

	//"github.com/kevinkimutai/metadata/internal/adapter/db/db"

	"github.com/kevinkimutai/metadata/internal/app/core/domain"
	"github.com/kevinkimutai/metadata/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

// Movies
func (a *Application) CreateNewMovie(movie domain.Movie) (domain.Movie, error) {

	movie, err := a.db.CreateMovie(movie)

	return movie, err
}

func (a *Application) GetMovies(movieParams domain.MovieParams) (domain.FetchData, error) {

	data, err := a.db.GetAllMovies(movieParams)

	return data, err
}

func (a *Application) GetMovie(movieID int64) (*domain.Movie, error) {
	//TODO:HANDLE ERRORS
	movie, err := a.db.GetMovieById(movieID)

	return movie, err
}

// Ratings
func (a *Application) CreateNewRating(rating domain.Rating) (domain.Rating, error) {

	rating, err := a.db.CreateRating(rating)

	return rating, err
}
