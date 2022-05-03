package api

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/token"
	"github.com/ndtai772/MyBookListBackend/util"
)

type Server struct {
	store             *db.Store
	router            *gin.Engine
	tokenMaker        token.Maker
	meilisearchClient *meilisearch.Client
}

const (
	DOMAIN_NAME = "api.mybooklist.ndtai.me"
)

func NewServer(store *db.Store) *Server {
	tokenMaker, err := token.NewJWTMaker(util.RandomString(32))
	if err != nil {
		panic(fmt.Errorf("cannot create JWT maker %w", err))
	}

	meilisearchClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://127.0.0.1:7700",
	})

	server := &Server{store: store, tokenMaker: tokenMaker, meilisearchClient: meilisearchClient}
	server.setupRouter()
	server.indexBooks()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	publicRoutes := router.Group("/")
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// Static
	publicRoutes.StaticFile("/", fmt.Sprintf("/var/www/%s/api-doc.html", DOMAIN_NAME))
	publicRoutes.StaticFile("/api-schema.yaml", fmt.Sprintf("/var/www/%s/api-schema.yaml", DOMAIN_NAME))

	//Auth
	publicRoutes.POST("/auth/login", server.login)
	publicRoutes.POST("/auth/refresh", server.renewAccessToken)

	// Accounts
	publicRoutes.POST("/accounts", server.createAccount)
	publicRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.PATCH("/accounts/:id", unimplemented("update account"))
	authRoutes.GET("/accounts/:id/bookmarks", server.listPersonalBookmarks)
	authRoutes.GET("/accounts/:id/rates", unimplemented("list personal rates"))
	authRoutes.GET("/accounts/:id/comments", unimplemented("list personal comments"))

	// Books
	publicRoutes.GET("/books", server.listBooks)
	publicRoutes.GET("/search", server.searchBooks)
	authRoutes.POST("/books", unimplemented("add book"))
	publicRoutes.GET("/books/:id", server.getBook)
	authRoutes.PATCH("/books/:id", unimplemented("update book"))
	authRoutes.DELETE("/books/:id", unimplemented("remove a book"))
	publicRoutes.GET("/books/:id/rates", unimplemented("list a book's rates"))
	publicRoutes.GET("/books/:id/comments", server.listBookComments)

	// Categories
	publicRoutes.GET("/categories", server.listCategories)
	authRoutes.POST("/categories", server.createCategory)
	publicRoutes.GET("/categories/:id", unimplemented("get category by id"))
	authRoutes.PATCH("/categories/:id", unimplemented("update category info"))
	authRoutes.DELETE("/categories/:id", unimplemented("delete a category"))
	publicRoutes.GET("/categories/:id/books", server.listBooksByCategory)

	// Bookmarks
	authRoutes.POST("/bookmarks", server.createBookmark)
	authRoutes.DELETE("/bookmarks/:id", server.deleteBookmark)

	// Rates
	authRoutes.POST("/rates", server.createRate)
	// authRoutes.GET("/rates/:id", unimplemented("get rate by id"))
	authRoutes.PATCH("/rates/:id", server.updateRate)
	authRoutes.DELETE("/rates/:id", unimplemented("delete a rate"))

	// Comments
	authRoutes.POST("/comments", server.createComment)
	authRoutes.DELETE("/comments/:id", server.deleteComment)

	server.router = router
}

func unimplemented(msg string) func(ctx *gin.Context) {
	log.Println("Unimplemented handler function: " + msg)
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"message": "unimplemented",
		})
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) indexBooks() {
	server.meilisearchClient.Index("books").DeleteAllDocuments()
	var wg sync.WaitGroup
	lastID := math.MaxInt32
	sem := make(chan int, 30)
	for {
		books, err := server.store.ListBooks(context.Background(), db.ListBooksParams{
			Limit:  1000,
			LastID: int32(lastID - 1),
		})
		if err != nil {
			log.Println("db error")
			break
		}
		if len(books) < 1 {
			break
		}
		wg.Add(1)
		sem <- 1
		go func(doc interface{}) {
			defer wg.Done()
			if _, err := server.meilisearchClient.Index("books").AddDocuments(doc); err != nil {
				panic(err)
			}
			<-sem
		}(books)

		lastID = int(books[len(books)-1].ID)
	}
	wg.Wait()
}