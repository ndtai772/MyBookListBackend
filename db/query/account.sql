-- name: CreateAccount :one
INSERT INTO accounts (
    name,
    email,
    hashed_password,
    avatar_url,
    is_admin
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByEmail :one
SELECT *
FROM accounts
WHERE email = $1 LIMIT 1;

-- -- name: UpdateAccount :one
-- UPDATE accounts
-- SET encoded_hash = $2
-- WHERE id = $1
-- RETURNING *;
