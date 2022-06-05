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

-- name: CheckRate :one
SELECT *
FROM rates
WHERE book_id = $1 AND created_by = $2
LIMIT 1;

-- name: UpdateRate :one
UPDATE rates
SET rate_value = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRate :exec
DELETE FROM rates
WHERE id = $1;

-- name: GetPersonalRateOfABook :one
SELECT *
FROM rates
WHERE created_by = $1 AND book_id = $2;
