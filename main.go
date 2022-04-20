package main

import (
	"database/sql"
	"log"

	"github.com/ndtai772/MyBookListBackend/api"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://dev:123@localhost:5432/my_book_list?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	server.Start(serverAddress)
}
