package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/meilisearch/meilisearch-go"
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

	var req struct {
		Limit  int `form:"limit" binding:"required"`
		Offset int `form:"offset"`
	}

	if err := ctx.ShouldBindWith(&req, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	books, err := server.bookIndex.Search("", &meilisearch.SearchRequest{
		Limit:  int64(req.Limit),
		Offset: int64(req.Offset),
		Filter: "categories = \"" + server.categoryIndex[int(id)] + "\"",
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":   books.Hits,
		"offset": books.Offset,
		"limit":  books.Limit,
	})
}
