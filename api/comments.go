package api

import (
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

	ctx.JSON(http.StatusOK, comment)
}

func (server *Server) deleteComment(ctx *gin.Context) {
	// TODO: verify ownership of the comment
	commentId, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteComment(ctx, commentId); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
