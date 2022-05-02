package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)

func (server *Server) listCategories(ctx *gin.Context) {
	categories, err := server.store.ListCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (server *Server) createCategory(ctx *gin.Context) {
	var createCategoryForm struct {
		Name        string `form:"name,omitempty"`
		Description string `form:"description,omitempty"`
	}

	if err := ctx.ShouldBindWith(&createCategoryForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createCategoryParams := db.CreateCategoryParams{
		Name:        createCategoryForm.Name,
		Description: createCategoryForm.Description,
	}

	category, err := server.store.CreateCategory(ctx, createCategoryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) listBooksByCategory(ctx *gin.Context) {
	id, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	page_size, last_id, err := parsePaginateQuery(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	listParams := db.ListBooksByCategoryIdParams{
		Limit:      page_size,
		LastID:     last_id,
		CategoryID: id,
	}

	books, err := server.store.ListBooksByCategoryId(ctx, listParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId := -1
	if len(books) > 0 {
		lastId = int(books[len(books)-1].ID)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":      books,
		"page_size": page_size,
		"last_id":   lastId,
	})
}
