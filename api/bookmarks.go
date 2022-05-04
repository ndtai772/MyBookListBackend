package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/token"
)

func (server *Server) createBookmark(ctx *gin.Context) {
	var createBookmarkForm struct {
		BookId int32 `form:"book_id" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&createBookmarkForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	createParams := db.CreateBookmarkParams{
		BookID:    createBookmarkForm.BookId,
		CreatedBy: authPayload.AccountID,
	}

	bookmark, err := server.store.CreateBookmark(ctx, createParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	server.updateBookIndex(bookmark.BookID)

	ctx.JSON(http.StatusOK, bookmark)
}

func (server *Server) deleteBookmark(ctx *gin.Context) {
	// TODO: verify ownership of the bookmark
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
