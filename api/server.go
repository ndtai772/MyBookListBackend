package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/token"
	"github.com/ndtai772/MyBookListBackend/util"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(store *db.Store) *Server {
	tokenMaker, err := token.NewJWTMaker(util.RandomString(32))
	if err != nil {
		panic(fmt.Errorf("cannot create JWT maker %w", err))
	}

	server := &Server{store: store, tokenMaker: tokenMaker}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	publicRoutes := router.Group("/")
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	//Auth
	publicRoutes.POST("/auth/login", server.login)

	// Accounts
	publicRoutes.POST("/accounts", server.createAccount)
	publicRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.PATCH("/accounts/:id", unimplemented("update account"))
	authRoutes.GET("/accounts/:id/bookmarks", server.listPersonalBookmarks)
	authRoutes.GET("/accounts/:id/rates", unimplemented("list personal rates"))
	authRoutes.GET("/accounts/:id/comments", unimplemented("list personal comments"))

	// Books
	publicRoutes.GET("/books", server.listBooks)
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
