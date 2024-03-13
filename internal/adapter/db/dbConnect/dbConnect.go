package dbconnect

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/kevinkimutai/metadata/internal/adapter/db/db"
)

type DBAdapter struct {
	ctx     context.Context
	queries *db.Queries
}

func NewDB(DBUrl string) *DBAdapter {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, DBUrl)
	if err != nil {
		log.Fatal("Error connecting to db:%w", err)
	}
	// defer conn.Close(ctx)

	queries := db.New(conn)

	return &DBAdapter{ctx: ctx, queries: queries}

}
func (a *DBAdapter) CreateMovie(movie db.Movie) (db.Movie, error) {
	//Map to CreateMovieParamsStruct
	movieParams := db.CreateMovieParams{
		Title:       movie.Title,
		Director:    movie.Director,
		Description: movie.Description,
	}

	movie, err := a.queries.CreateMovie(a.ctx, movieParams)

	return movie, err

}

func (a *DBAdapter) GetMovieById(movieID int64) (db.Movie, error) {
	movie, err := a.queries.GetMovie(a.ctx, movieID)

	if err != nil {
		//Dismiss No Movie With ID Error
		if err.Error() == "no rows in result set" {
			movie = db.Movie{
				ID: movie.ID,
			}
			err = nil

			return movie, err

		}
		return movie, err

	}

	return movie, nil

}

// func (a *DBAdapter) GetAllMovies(movieID int64) {

// }

// func (a *DBAdapter) CreateMovie(movieID int64) {

// }

// func (a *DBAdapter) UpdateMovie(movieID int64) {

// }

// func (a *DBAdapter) DeleteMovie(movieID int64) {

// }
