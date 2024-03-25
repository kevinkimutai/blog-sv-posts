// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (
  title, description, director
) VALUES (
  $1, $2, $3
)
RETURNING id, title, description, director, created_at
`

type CreateMovieParams struct {
	Title       string
	Description string
	Director    string
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRow(ctx, createMovie, arg.Title, arg.Description, arg.Director)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Director,
		&i.CreatedAt,
	)
	return i, err
}

const createRating = `-- name: CreateRating :one
INSERT INTO ratings (
  movie_id, rating
) VALUES (
  $1, $2
)
RETURNING id, movie_id, rating, created_at
`

type CreateRatingParams struct {
	MovieID int64
	Rating  pgtype.Numeric
}

func (q *Queries) CreateRating(ctx context.Context, arg CreateRatingParams) (Rating, error) {
	row := q.db.QueryRow(ctx, createRating, arg.MovieID, arg.Rating)
	var i Rating
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.Rating,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE FROM movies
WHERE movies.id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteMovie, id)
	return err
}

const getMovie = `-- name: GetMovie :one
SELECT movies.id, movies.title, movies.description, movies.director, movies.created_at,COALESCE(AVG(ratings.rating), 0.0) AS average_rating
FROM movies
LEFT JOIN ratings 
ON movies.id = ratings.movie_id
WHERE movies.id = $1 
GROUP BY movies.id
LIMIT 1
`

type GetMovieRow struct {
	ID            int64
	Title         string
	Description   string
	Director      string
	CreatedAt     pgtype.Timestamptz
	AverageRating interface{}
}

func (q *Queries) GetMovie(ctx context.Context, id int64) (GetMovieRow, error) {
	row := q.db.QueryRow(ctx, getMovie, id)
	var i GetMovieRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Director,
		&i.CreatedAt,
		&i.AverageRating,
	)
	return i, err
}

const listMovies = `-- name: ListMovies :many
SELECT movies.id, movies.title, movies.description, movies.director, movies.created_at,AVG(ratings.rating) AS average_rating
FROM movies
LEFT JOIN ratings 
ON movies.id = ratings.movie_id
`

type ListMoviesRow struct {
	ID            int64
	Title         string
	Description   string
	Director      string
	CreatedAt     pgtype.Timestamptz
	AverageRating float64
}

func (q *Queries) ListMovies(ctx context.Context) ([]ListMoviesRow, error) {
	rows, err := q.db.Query(ctx, listMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMoviesRow
	for rows.Next() {
		var i ListMoviesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Director,
			&i.CreatedAt,
			&i.AverageRating,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :exec
UPDATE movies
  set title = $1,
  description = $2,
  director= $3
WHERE movies.id = $1
`

type UpdateMovieParams struct {
	Title       string
	Description string
	Director    string
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) error {
	_, err := q.db.Exec(ctx, updateMovie, arg.Title, arg.Description, arg.Director)
	return err
}
