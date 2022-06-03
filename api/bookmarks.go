package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/token"
)

func (server *Server) createBookmark(ctx *gin.Context) {
	var createBookmarkForm struct {
		BookId int32 `form:"book_id" binding:"required"`
		Type   int32 `form:"bookmark_type" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&createBookmarkForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	createParams := db.CreateBookmarkParams{
		BookID:    createBookmarkForm.BookId,
		CreatedBy: authPayload.AccountID,
		Type:      createBookmarkForm.Type,
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
	id, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	bookmarks, err := server.store.GetBookmark(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if bookmarks.CreatedBy != authPayload.AccountID {
		ctx.JSON(http.StatusForbidden, errorResponse(errors.New("you don't own this bookmark")))
		return
	}

	if err := server.store.DeleteBookmark(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) updateBookmarkType(ctx *gin.Context) {
	id, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	bookmarks, err := server.store.GetBookmark(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if bookmarks.CreatedBy != authPayload.AccountID {
		ctx.JSON(http.StatusForbidden, errorResponse(errors.New("you don't own this bookmark")))
		return
	}

	var updateBookmarkForm struct {
		Type   int32 `form:"bookmark_type" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&updateBookmarkForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updateBookmarkTypeParams := db.UpdateBookmarkTypeParams {
		NewBookmarkType: updateBookmarkForm.Type,
	}

	updatedBookmark, err := server.store.UpdateBookmarkType(ctx, updateBookmarkTypeParams);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedBookmark)
}

func (server *Server) listPersonalBookmarks(ctx *gin.Context) {
	id, err := parseIdUri(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.AccountID != id {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	bookmarks, err := server.store.ListBookmarkedBooksByAccountId(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookmarks,
	})
}
