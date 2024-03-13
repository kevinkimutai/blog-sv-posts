package ports

import "github.com/kevinkimutai/metadata/internal/adapter/db/db"

type DBPort interface {
	GetMovieById(movieID int64) (db.Movie, error)
	CreateMovie(movie db.Movie) (db.Movie, error)
}
