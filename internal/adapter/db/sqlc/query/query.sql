-- name: GetMovie :one
SELECT movies.*,COALESCE(AVG(ratings.rating), 0.0) AS average_rating
FROM movies
LEFT JOIN ratings 
ON movies.id = ratings.movie_id
WHERE movies.id = $1 
GROUP BY movies.id
LIMIT 1;

-- name: ListMovies :many
SELECT movies.*,AVG(ratings.rating) AS average_rating
FROM movies
LEFT JOIN ratings 
ON movies.id = ratings.movie_id;


-- name: CreateMovie :one
INSERT INTO movies (
  title, description, director
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateMovie :exec
UPDATE movies
  set title = $1,
  description = $2,
  director= $3
WHERE movies.id = $1;

-- name: DeleteMovie :exec
DELETE FROM movies
WHERE movies.id = $1;

-- name: CreateRating :one
INSERT INTO ratings (
  movie_id, rating
) VALUES (
  $1, $2
)
RETURNING *;