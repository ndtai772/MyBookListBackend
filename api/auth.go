package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ndtai772/MyBookListBackend/util"
)

var accessTokenDuration = time.Duration.Minutes(15)
var refreshTokenDuration = time.Duration.Hours(7 * 24)

func (server *Server) login(ctx *gin.Context) {
	var loginReq struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindWith(&loginReq, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccountByEmail(ctx, loginReq.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	if err := util.CheckPassword(loginReq.Password, account.HashedPassword); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		account.ID,
		time.Duration(accessTokenDuration),
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, _, err := server.tokenMaker.CreateToken(
		account.ID,
		time.Duration(refreshTokenDuration),
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"access_token": accessToken,
		"refresh_token": refreshToken,
	})
}
