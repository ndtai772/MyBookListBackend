-- name: CreateRate :one
INSERT INTO rates (
    book_id,
    created_by,
    rate_value
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetRate :one
SELECT *
FROM rates
WHERE id = $1
LIMIT 1;

-- name: UpdateRate :one
UPDATE rates
SET rate_value = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRate :exec
DELETE FROM rates
WHERE id = $1;

-- name: ListRatesByAccountId :many
SELECT *
FROM rates
WHERE created_by = $3
LIMIT $1
OFFSET $2;
