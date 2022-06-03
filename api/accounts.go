package api

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
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
		AvatarUrl: account.AvatarUrl,
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
		AvatarUrl:      fmt.Sprintf("https://ui-avatars.com/api/?name=%s", url.QueryEscape(reqForm.Name)),
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, toAccountRes(account))
}

func (server *Server) getAccountInfo(ctx *gin.Context) {
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
