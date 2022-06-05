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

-- name: GetBookBrief :one
SELECT id, title, author, publisher, cover_url, categories, comment_count, bookmark_count, rate_count, rate_avg, pages
FROM book_detail
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT id, title, author, publisher, cover_url, categories, comment_count, bookmark_count, rate_count, rate_avg, pages
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

-- name: ListBooksByCategoryId :many
SELECT books.*
FROM books
LEFT JOIN book_category
ON  book_category.category_id = $2
    AND book_category.book_id = books.id
WHERE NOT id > @last_id
ORDER BY books.id DESC
LIMIT $1;