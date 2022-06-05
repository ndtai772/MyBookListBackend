-- name: CreateBookmark :one
INSERT INTO bookmarks (
    book_id,
    type,
    created_by
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetBookmark :one
SELECT *
FROM bookmarks
WHERE id = $1;

-- name: CheckBookmark :one
SELECT *
FROM bookmarks
WHERE book_id = $1 AND created_by = $2;

-- name: UpdateBookmarkType :one
UPDATE bookmarks
SET type = @new_bookmark_type
WHERE id = @id
RETURNING *;

-- name: DeleteBookmark :exec
DELETE FROM bookmarks
WHERE id = $1;

-- name: ListBookmarkedBooksByAccountId :many
SELECT bd.id,
       title,
       author,
       description,
       year,
       language,
       publisher,
       pages,
       cover_url,
       categories,
       comment_count,
       bookmark_count,
       rate_count,
       rate_avg,
       bookmarks.id   as bookmark_id,
       bookmarks.type as bookmark_type
FROM bookmarks
JOIN book_detail bd on bookmarks.book_id = bd.id
WHERE bookmarks.created_by = $1
ORDER BY bookmarks.id DESC;
