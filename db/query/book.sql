-- name: CreateBook :one
INSERT INTO books (
    title,
    author,
    description,
    year,
    language,
    publisher,
    pages
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetBook :one
SELECT *
FROM book_detail
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT id, title, author, language, cover_url, created_at, categories, comment_count, bookmark_count, rate_count, rate_sum
FROM book_detail
WHERE NOT id > @last_id
ORDER BY id DESC
LIMIT $1;

-- -- name: UpdateBook :one
-- UPDATE books
-- SET title = $2,
--     author = $3,
--     description = $4
-- WHERE id = $1
-- RETURNING *;

-- -- name: DeleteBook :exec
-- DELETE FROM books
-- WHERE id = $1;
