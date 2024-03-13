package application

import (

	//"github.com/kevinkimutai/metadata/internal/adapter/db/db"
	"github.com/kevinkimutai/metadata/internal/adapter/db/db"
	"github.com/kevinkimutai/metadata/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) CreateNewMovie(movie db.Movie) (db.Movie, error) {

	movie, err := a.db.CreateMovie(movie)

	return movie, err
}

func (a *Application) GetMovie(movieID int64) (db.Movie, error) {
	//TODO:HANDLE ERRORS
	movie, err := a.db.GetMovieById(movieID)

	return movie, err
}
