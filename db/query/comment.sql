-- name: CreateComment :one
INSERT INTO comments (
    content,
    book_id,
    created_by
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetComment :one
SELECT *
FROM comments
WHERE id = $1 LIMIT 1;

-- name: UpdateComment :one
UPDATE comments
SET content = $2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;

-- name: ListCommentsByBookId :many
SELECT comments.*,
    COALESCE(a.name, '')        as username,
    COALESCE(a.avatar_url, '')  as avatar_url,
    COALESCE(a.is_admin, false) as is_admin
FROM comments
    LEFT JOIN accounts a
        ON a.id = comments.created_by
WHERE book_id = @book_id AND NOT comments.id > @last_id
ORDER BY comments.id DESC
LIMIT @page_size;

-- name: ListCommentsByAccountId :many
SELECT comments.*,
    b.title,
    b.cover_url
FROM comments
    JOIN books b on b.id = comments.book_id
WHERE created_by = @user_id AND NOT comments.id > @last_id
ORDER BY comments.id DESC
LIMIT @page_size;