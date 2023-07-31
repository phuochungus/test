-- name: GetAuthor :one
SELECT *
FROM authors
WHERE id = $1
LIMIT 1;

-- name: GetAuthorByName :one
SELECT *
FROM authors
WHERE name = $1
LIMIT 1;

-- name: ListAuthors :many
SELECT *
FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: UpdateAuthor :one
UPDATE authors
SET name = @name::TEXT
WHERE id = @id::BIGINT
RETURNING *;