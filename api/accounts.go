package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)

func (server *Server) createAccount(ctx *gin.Context) {
	var createAccountRequest struct {
		Username string `form:"username" binding:"required,alphanum"`
		Password string `form:"password" binding:"required,min=6"`
		Email    string `form:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindWith(&createAccountRequest, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Username:    createAccountRequest.Username,
		Email:       createAccountRequest.Email,
		EncodedHash: createAccountRequest.Password,
		IsAdmin:     false,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) getAccount(ctx *gin.Context) {
	var getAccountRequest struct {
		Id int32 `uri:"id" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindUri(&getAccountRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccount(ctx, getAccountRequest.Id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) listPersonalBookmarks(ctx *gin.Context) {
	id, err := parseIdUri(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	limit, offset, err := parsePaginateQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	bookmarks, err := server.store.ListBookmarksByAccountId(ctx, db.ListBookmarksByAccountIdParams{
		Limit:     limit,
		Offset:    offset,
		CreatedBy: id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"data": bookmarks,
		"next_index": offset + int32(len(bookmarks)),
	})
}

func (server *Server) listPersonalRates(ctx * gin.Context) {
	id, err := parseIdUri(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	limit, offset, err := parsePaginateQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	rates, err := server.store.ListRatesByAccountId(ctx, db.ListRatesByAccountIdParams{
		Limit:     limit,
		Offset:    offset,
		CreatedBy: id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"data": rates,
		"next_index": offset + int32(len(rates)),
	})
}
