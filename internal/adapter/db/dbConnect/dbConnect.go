package dbconnect

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/kevinkimutai/metadata/internal/adapter/db/db"
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
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

// Movies
func (a *DBAdapter) CreateMovie(movie domain.Movie) (domain.Movie, error) {

	//Map to CreateMovieParamsStruct
	movieParams := db.CreateMovieParams{
		Title:       movie.Title,
		Director:    movie.Director,
		Description: movie.Description,
	}

	mov, err := a.queries.CreateMovie(a.ctx, movieParams)

	//Map structs
	data := domain.Movie{
		ID:          mov.ID,
		Description: mov.Description,
		Title:       mov.Title,
		Director:    mov.Director,
		CreatedAt:   mov.CreatedAt.Time,
	}

	return data, err

}

func (a *DBAdapter) GetMovieById(movieID int64) (*domain.Movie, error) {
	movie, err := a.queries.GetMovie(a.ctx, movieID)
	fmt.Println("Movie : ", movie)

	avgRating, ok := interfaceToFloat64(movie.AverageRating)
	if ok != nil {
		return &domain.Movie{}, err
	}

	//Map Struct
	data := &domain.Movie{
		ID:            movie.ID,
		Title:         movie.Title,
		Description:   movie.Description,
		Director:      movie.Director,
		AverageRating: avgRating,
		CreatedAt:     movie.CreatedAt.Time,
	}

	if err != nil {
		//Dismiss No Movie With ID Error
		if err.Error() == "no rows in result set" {
			err = nil
			return data, err
		}
		return data, err

	}

	return data, nil

}

func interfaceToFloat64(value interface{}) (float64, error) {
	// Check if the value is already a float64
	if f, ok := value.(float64); ok {
		return f, nil
	}

	// Check if the value is of type pgtype.Numeric
	if numeric, ok := value.(pgtype.Numeric); ok {
		fval, err := numeric.Value()
		if err != nil {
			return 0, err
		}

		//Convert To Float64
		var floatVal float64
		if strVal, ok := fval.(string); ok {
			floatVal, err = strconv.ParseFloat(strVal, 64)
			if err != nil {
				return 0, err
			}
		} else {
			// Handle the case where fval is not a string
			return 0, fmt.Errorf("value is not a string")
		}

		return floatVal, nil
	}

	// If not float64 or pgtype.Numeric, return error
	return 0, errors.New("value cannot be converted to float64")
}

// Ratings
func (a *DBAdapter) CreateRating(rating domain.Rating) (domain.Rating, error) {
	var numeric pgtype.Numeric

	numeric.Scan(rating.Rating)
	//Map Struct
	ratings := db.CreateRatingParams{
		MovieID: rating.MovieID,
		Rating:  numeric,
	}

	dbRating, err := a.queries.CreateRating(a.ctx, ratings)

	//Map Struct
	data := domain.Rating{
		ID:        dbRating.ID,
		MovieID:   dbRating.MovieID,
		Rating:    rating.Rating,
		CreatedAt: dbRating.CreatedAt.Time,
	}

	return data, err

}
