// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: book.sql

package db

import (
	"context"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (
    title,
    author,
    description
) VALUES (
    $1, $2, $3
) RETURNING id, title, author, description, modified_at, created_at
`

type CreateBookParams struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook, arg.Title, arg.Author, arg.Description)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Description,
		&i.ModifiedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const getBook = `-- name: GetBook :one
SELECT id, title, author, description, modified_at, created_at
FROM books
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Description,
		&i.ModifiedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT id, title, author, description, modified_at, created_at
FROM books
LIMIT $1
OFFSET $2
`

type ListBooksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.Description,
			&i.ModifiedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBook = `-- name: UpdateBook :one
UPDATE books
SET title = $2,
    author = $3,
    description = $4
WHERE id = $1
RETURNING id, title, author, description, modified_at, created_at
`

type UpdateBookParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBook,
		arg.ID,
		arg.Title,
		arg.Author,
		arg.Description,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Description,
		&i.ModifiedAt,
		&i.CreatedAt,
	)
	return i, err
}
