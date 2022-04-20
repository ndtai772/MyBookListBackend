package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	publicRoutes := router.Group("/")
	authRoutes := router.Group("/")

	// Accounts
	publicRoutes.POST("/accounts", server.createAccount)
	publicRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.PATCH("/accounts/:id", unimplemented("update account info"))
	authRoutes.DELETE("/accounts/:id", unimplemented("inactive account"))
	authRoutes.GET("/accounts/:id/feedbacks", unimplemented("get personal feedbacks"))
	authRoutes.GET("/accounts/:id/bookmarks", server.listPersonalBookmarks)
	authRoutes.GET("/accounts/:id/rates", server.listPersonalRates)

	// Feedbacks
	authRoutes.GET("/feedbacks", unimplemented("list all user feedbacks"))
	authRoutes.POST("/feedbacks", unimplemented("create feedback"))
	authRoutes.GET("/feedbacks/:id", unimplemented("get feedback by id"))
	authRoutes.PATCH("/feedbacks/:id", unimplemented("update feedback"))
	authRoutes.DELETE("/feedbacks/:id", unimplemented("delete a feedback"))

	// Books
	publicRoutes.GET("/books", server.listBooks)
	authRoutes.POST("/books", server.createBook)
	authRoutes.GET("/books/:id", server.getBook)
	authRoutes.PATCH("/books/:id", server.updateBook)
	authRoutes.DELETE("/books/:id", server.deleteBook)
	authRoutes.GET("/books/:id/rates", server.listBookRates)
	authRoutes.GET("/books/:id/comments", server.listBookComments)

	// Categories
	authRoutes.GET("/categories", server.listCategories)
	authRoutes.POST("/categories", server.createCategory)
	// authRoutes.GET("/categories/:id", unimplemented("get category by id"))
	authRoutes.PATCH("/categories/:id", unimplemented("update category info"))
	authRoutes.DELETE("/categories/:id", unimplemented("delete a category"))
	authRoutes.GET("/categories/:id/books", server.listBooksByCategory)

	// Bookmarks
	authRoutes.POST("/bookmarks", server.createBookmark)
	authRoutes.DELETE("/bookmarks/:id", server.deleteBookmark)

	// Rates
	authRoutes.POST("/rates", server.createRate)
	// authRoutes.GET("/rates/:id", unimplemented("get rate by id"))
	authRoutes.PATCH("/rates/:id", server.updateRate)
	authRoutes.DELETE("/rates/:id", unimplemented("delete a rate"))

	// Comments
	authRoutes.POST("/comments", unimplemented("create comment"))
	authRoutes.GET("/comments/:id", unimplemented("get rate by id"))
	authRoutes.DELETE("/comments/:id", unimplemented("delete a rate"))

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
