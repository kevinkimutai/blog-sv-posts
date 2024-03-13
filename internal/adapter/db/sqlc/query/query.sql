-- name: GetMovie :one
SELECT * FROM movies
WHERE id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies
ORDER BY title;

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
WHERE id = $1;

-- name: DeleteMovie :exec
DELETE FROM movies
WHERE id = $1;