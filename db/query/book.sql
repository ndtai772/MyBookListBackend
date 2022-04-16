-- name: CreateBook :one
INSERT INTO books (
    title,
    author,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetBook :one
SELECT *
FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT *
FROM books
LIMIT $1
OFFSET $2;

-- name: UpdateBook :one
UPDATE books
SET title = $2,
    author = $3,
    description = $4
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;
