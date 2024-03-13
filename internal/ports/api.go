package ports

import "github.com/kevinkimutai/metadata/internal/adapter/db/db"

type APIPort interface {
	GetMovie(movieID int64) (db.Movie, error)
	CreateNewMovie(movie db.Movie) (db.Movie, error)
}
