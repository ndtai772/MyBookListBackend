-- name: CreateBookCategory :one
INSERT INTO book_category (
    book_id,
    category_id
) VALUES (
    $1, $2
) RETURNING *;

-- -- name: DeleteBookCategory :exec
-- DELETE FROM book_category
-- WHERE book_id = $1 AND category_id = $2;
