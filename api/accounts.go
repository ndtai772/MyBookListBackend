package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/token"
	"github.com/ndtai772/MyBookListBackend/util"
)

type AccountRes struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarUrl string    `json:"avatar_url"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

func toAccountRes(account db.Account) AccountRes {
	return AccountRes{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		AvatarUrl: fmt.Sprintf("%s%s", "{{baseUrl}}/resources/images", account.AvatarUrl),
		IsAdmin:   account.IsAdmin,
		CreatedAt: account.CreatedAt,
	}
}

func (server *Server) createAccount(ctx *gin.Context) {
	var reqForm struct {
		Name     string `form:"name" binding:"required"`
		Password string `form:"password" binding:"required,min=6"`
		Email    string `form:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindWith(&reqForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPw, err := util.HashPassword(reqForm.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Name:           reqForm.Name,
		Email:          reqForm.Email,
		HashedPassword: hashedPw,
		IsAdmin:        false,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, toAccountRes(account))
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

	ctx.JSON(http.StatusOK, toAccountRes(account))
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
		"data":      bookmarks,
	})
}

// func (server *Server) listPersonalRates(ctx *gin.Context) {
// 	id, err := parseIdUri(ctx)

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 	}

// 	page_size, last_id, err := parsePaginateQuery(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 	}

// 	rates, err := server.store.ListRatesByAccountId(ctx, db.ListRatesByAccountIdParams{
// 		Limit:     page_size,
// 		last_id:   last_id,
// 		CreatedBy: id,
// 	})

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data":       rates,
// 		"next_index": offset + int32(len(rates)),
// 	})
// }

// func (server *Server) listPersonalComments(ctx *gin.Context) {
// 	accountId, err := parseIdUri(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	page_size, last_id, err := parsePaginateQuery(ctx)

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	comments, err := server.store.ListCommentsByAccoutId(ctx, db.ListCommentsByBookIdParams{
// 		Limit:     page_size,
// 		Offset:    offset,
// 		CreatedBy: accountId,
// 	})

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data":       comments,
// 		"next_index": offset + int32(len(comments)),
// 	})
// }
