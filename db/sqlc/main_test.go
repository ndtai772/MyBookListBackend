package db

import "testing"

const (
	dbDriver = "postgres"
	dbSource = "postgresql://dev:123@localhost:5432/my_book_list?sslmode=disable"
)


func TestMain(m *testing.M) {
	
}