package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/token"
)

func (server *Server) createComment(ctx *gin.Context) {
	var createCommentForm struct {
		BookId  int32  `form:"book_id" binding:"required"`
		Content string `form:"content" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&createCommentForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	createCommentParams := db.CreateCommentParams{
		Content:   createCommentForm.Content,
		BookID:    createCommentForm.BookId,
		CreatedBy: authPayload.AccountID,
	}

	comment, err := server.store.CreateComment(ctx, createCommentParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	server.updateBookIndex(comment.BookID)

	ctx.JSON(http.StatusOK, comment)
}

func (server *Server) deleteComment(ctx *gin.Context) {
	commentId, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comment, err := server.store.GetBookmark(ctx, commentId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if comment.CreatedBy != authPayload.AccountID {
		ctx.JSON(http.StatusForbidden, errorResponse(errors.New("you don't own this comment")))
		return
	}

	if err := server.store.DeleteComment(ctx, commentId); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	server.updateBookIndex(comment.BookID)

	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) listPersonalComments(ctx *gin.Context) {
	accountId, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.AccountID != accountId {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	page_size, last_id, err := parsePaginateQuery(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comments, err := server.store.ListCommentsByAccountId(ctx, db.ListCommentsByAccountIdParams{
		UserID:   accountId,
		PageSize: page_size,
		LastID:   last_id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastID := -1

	if len(comments) > 0 {
		lastID = int(comments[len(comments)-1].ID)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    comments,
		"last_id": lastID,
	})
}
