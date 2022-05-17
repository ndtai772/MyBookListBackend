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
FROM comment_detail
WHERE id = $1 LIMIT 1;

-- name: UpdateComment :one
UPDATE comments
SET content = $2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;

-- name: ListCommentsByBookId :many
SELECT *
FROM comment_detail
WHERE book_id = $2 AND NOT id > @last_id
ORDER BY id DESC
LIMIT $1;
