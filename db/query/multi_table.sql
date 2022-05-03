-- name: ListBooksByCategoryId :many
SELECT books.*
FROM books
LEFT JOIN book_category
ON  book_category.category_id = $2
    AND book_category.book_id = books.id
WHERE NOT id > @last_id
ORDER BY books.id DESC
LIMIT $1;

-- name: ListBookmarkedBooksByAccountId :many
SELECT books.id, books.title, books.author, books.language, books.publisher, books.pages, books.cover_url
FROM bookmarks
JOIN books on books.id = bookmarks.book_id
WHERE bookmarks.created_by = $1
ORDER BY bookmarks.id DESC;
