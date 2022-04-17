-- name: CreateComment :one
INSERT INTO comments (
    content,
    book_id,
    created_by
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetComment :one
SELECT *
FROM comments
WHERE id = $1 LIMIT 1;

-- name: ListComments :many
SELECT *
FROM comments
LIMIT $1
OFFSET $2;

-- name: UpdateComment :one
UPDATE comments
SET content = $2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;