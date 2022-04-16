// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
)

type Account struct {
	ID          int32          `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	EncodedHash string         `json:"encoded_hash"`
	AvatarUri   sql.NullString `json:"avatar_uri"`
	IsAdmin     bool           `json:"is_admin"`
	ModifiedAt  sql.NullTime   `json:"modified_at"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type Book struct {
	ID          int32        `json:"id"`
	Title       string       `json:"title"`
	Author      string       `json:"author"`
	Description string       `json:"description"`
	ModifiedAt  sql.NullTime `json:"modified_at"`
	CreatedAt   sql.NullTime `json:"created_at"`
}

type BookCategory struct {
	BookID     int32        `json:"book_id"`
	CategoryID int32        `json:"category_id"`
	CreatedAt  sql.NullTime `json:"created_at"`
}

type Bookmark struct {
	ID        int32        `json:"id"`
	BookID    int32        `json:"book_id"`
	CreatedBy int32        `json:"created_by"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Category struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	ModifiedAt  sql.NullTime   `json:"modified_at"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type Comment struct {
	ID         int32        `json:"id"`
	Content    string       `json:"content"`
	BookID     int32        `json:"book_id"`
	CreatedBy  int32        `json:"created_by"`
	ModifiedAt sql.NullTime `json:"modified_at"`
	CreatedAt  sql.NullTime `json:"created_at"`
}

type Feedback struct {
	ID           int32          `json:"id"`
	Content      string         `json:"content"`
	CreatedBy    int32          `json:"created_by"`
	IsViewed     bool           `json:"is_viewed"`
	IsProcessing bool           `json:"is_processing"`
	IsResolved   bool           `json:"is_resolved"`
	Message      sql.NullString `json:"message"`
	ModifiedAt   sql.NullTime   `json:"modified_at"`
	CreatedAt    sql.NullTime   `json:"created_at"`
}

type Rate struct {
	ID         int32        `json:"id"`
	BookID     int32        `json:"book_id"`
	CreatedBy  int32        `json:"created_by"`
	RateValue  int32        `json:"rate_value"`
	ModifiedAt sql.NullTime `json:"modified_at"`
	CreatedAt  sql.NullTime `json:"created_at"`
}
