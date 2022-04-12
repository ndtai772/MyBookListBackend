-- name: CreateBookmark :one
INSERT INTO bookmarks (
    book_id,
    created_by
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetBookmark :one
SELECT *
FROM bookmarks
WHERE id = $1 LIMIT 1;

-- name: DeleteBookmark :exec
DELETE FROM bookmarks
WHERE id = $1;
