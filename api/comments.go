package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)

func (server *Server) createComment(ctx *gin.Context) {
	var createCommentForm struct {
		BookId      int32  `json:"book_id,omitempty"`
		Content     string `json:"content,omitempty"`
		CreatedBy   int32  `json:"created_by,omitempty"`
	}

	if err := ctx.ShouldBindWith(&createCommentForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createCommentParams := db.CreateCommentParams{
		Content: createCommentForm.Content,
		BookID: createCommentForm.BookId,
		CreatedBy: createCommentForm.CreatedBy,
	}

	comment, err := server.store.CreateComment(ctx, createCommentParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (server *Server) deleteComment(ctx *gin.Context) {
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
