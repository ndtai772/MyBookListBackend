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
	authRoutes.GET("/accounts/:id/bookmarks", unimplemented("get personal bookmarks"))
	authRoutes.GET("/accounts/:id/rates", unimplemented("get personal rates"))

	// Feedbacks
	authRoutes.GET("/feedbacks", unimplemented("list all user feedbacks"))
	authRoutes.POST("/feedbacks", unimplemented("create feedback"))
	authRoutes.GET("/feedbacks/:id", unimplemented("get feedback by id"))
	authRoutes.PATCH("/feedbacks/:id", unimplemented("update feedback"))
	authRoutes.DELETE("/feedbacks/:id", unimplemented("delete a feedback"))

	// Books
	authRoutes.GET("/books", unimplemented("list books"))
	authRoutes.POST("/books", unimplemented("create book"))
	authRoutes.GET("/books/:id", unimplemented("get book by id"))
	authRoutes.PATCH("/books/:id", unimplemented("update book info"))
	authRoutes.DELETE("/books/:id", unimplemented("delete a book"))
	authRoutes.GET("/books/:id/rates", unimplemented("get rates of a book"))
	authRoutes.GET("/books/:id/comments", unimplemented("get comments of a book"))

	// Categories
	authRoutes.GET("/categories", unimplemented("list categories"))
	authRoutes.POST("/categories", unimplemented("create category"))
	authRoutes.GET("/categories/:id", unimplemented("get category by id"))
	authRoutes.PATCH("/categories/:id", unimplemented("update category info"))
	authRoutes.DELETE("/categories/:id", unimplemented("delete a category"))

	// Bookmarks
	authRoutes.POST("/bookmarks", unimplemented("create bookmark"))
	authRoutes.DELETE("/bookmarks/:id", unimplemented("delete a bookmark"))

	// Rates
	authRoutes.POST("/rates", unimplemented("create rate"))
	authRoutes.GET("/rates/:id", unimplemented("get rate by id"))
	authRoutes.PATCH("/rates/:id", unimplemented("update rate"))
	authRoutes.DELETE("/rates/:id", unimplemented("delete a rate"))

	// Comments
	authRoutes.POST("/comments", unimplemented("create comment"))
	authRoutes.GET("/comments/:id", unimplemented("get rate by id"))
	authRoutes.DELETE("/comments/:id", unimplemented("delete a rate"))

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
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
