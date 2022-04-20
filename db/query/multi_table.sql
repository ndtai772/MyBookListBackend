-- name: ListBooksByCategoryId :many
SELECT books.*
FROM books
LEFT JOIN book_category
ON  book_category.category_id = $3
    AND book_category.book_id = books.id
ORDER BY books.id
LIMIT $1
OFFSET $2;
