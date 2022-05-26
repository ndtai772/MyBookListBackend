package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/token"
)

type Server struct {
	store         *db.Store
	router        *gin.Engine
	tokenMaker    token.Maker
	bookIndex     *meilisearch.Index
	categoryIndex map[int]string
}

const (
	DOMAIN_NAME      = "api.mybooklist.ndtai.me"
	MEILISEARCH_HOST = "http://127.0.0.1:7700"
	TOKEN_SECRET     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXcVCJ9"
)

func NewServer(store *db.Store) *Server {
	tokenMaker, err := token.NewJWTMaker(TOKEN_SECRET)
	if err != nil {
		panic(fmt.Errorf("cannot create JWT maker %w", err))
	}

	meilisearchClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: MEILISEARCH_HOST,
	})

	bookIndex := meilisearchClient.Index("books")
	bookIndex.UpdateSettings(&meilisearch.Settings{
		SearchableAttributes: []string{
			"title",
			"author",
			"publisher",
			"categories",
		},
	})

	categories, err := store.ListCategories(context.Background())
	if err != nil {
		panic("couldn't get categories from db")
	}
	categoryIndex := map[int]string{}
	for _, category := range categories {
		categoryIndex[int(category.ID)] = category.Name
	}

	server := &Server{store: store, tokenMaker: tokenMaker, bookIndex: meilisearchClient.Index("books"), categoryIndex: categoryIndex}
	server.setupRouter()
	server.indexBooks()
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
