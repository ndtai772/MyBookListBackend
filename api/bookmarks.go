package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)

func (server *Server) createBookmark(ctx *gin.Context) {
	var createBookmarkForm struct {
		BookId    int32 `form:"book_id" binding:"required"`
		CreatedBy int32 `form:"created_by" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&createBookmarkForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createParams := db.CreateBookmarkParams{
		BookID:    createBookmarkForm.BookId,
		CreatedBy: createBookmarkForm.CreatedBy,
	}

	bookmark, err := server.store.CreateBookmark(ctx, createParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, bookmark)
}

func (server *Server) deleteBookmark(ctx *gin.Context) {
	id, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteBookmark(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
