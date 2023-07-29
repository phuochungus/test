-- name: GetAuthor :one
SELECT *
FROM authors
WHERE id = $1
LIMIT 1;

-- name: ListAuthors :many
SELECT *
FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name)
VALUES ($1)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: UpdateAuthor :one
UPDATE authors
SET name = coalesce(sqlc.narg('name'), name)
WHERE id = @id::BIGINT
RETURNING *;