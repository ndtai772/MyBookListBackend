-- name: CreateFeedback :one
INSERT INTO feedbacks (
    content,
    created_by
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetFeedback :one
SELECT *
FROM feedbacks
WHERE id = $1 LIMIT 1;

-- name: ListFeedbacks :many
SELECT *
FROM feedbacks
LIMIT $1
OFFSET $2;

-- name: UpdateFeedback :one
UPDATE feedbacks
SET content = $2
WHERE id = $1
RETURNING *;

-- name: DeleteFeedback :exec
DELETE FROM feedbacks
WHERE id = $1;
