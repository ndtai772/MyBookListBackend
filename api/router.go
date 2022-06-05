package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	publicRoutes.GET("/accounts/:id", server.getAccountInfo)
	authRoutes.PATCH("/accounts/:id", unimplemented("update account info"))
	authRoutes.PATCH("/accounts/:id/password", unimplemented("update password"))
	authRoutes.GET("/accounts/:id/bookmarks", server.listPersonalBookmarks)
	authRoutes.GET("/accounts/:id/rates", unimplemented("list personal rates"))
	authRoutes.GET("/accounts/:id/comments", server.listPersonalComments)

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
	authRoutes.GET("/bookmarks", server.checkBookmark)
	authRoutes.POST("/bookmarks", server.createBookmark)
	authRoutes.DELETE("/bookmarks/:id", server.deleteBookmark)
	authRoutes.PATCH("/bookmarks/:id", server.updateBookmarkType)

	// Rates
	authRoutes.GET("/rates", server.checkRate)
	authRoutes.POST("/rates", server.createRate)
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
