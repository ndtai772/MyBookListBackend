-- name: CreateAccount :one
INSERT INTO accounts (
    id,
    username,
    email,
    is_admin
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT *
FROM accounts
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET encoded_hash = $2
WHERE id = $1
RETURNING *;
