package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ndtai772/MyBookListBackend/util"
)

var accessTokenDuration = time.Minute * 15
var refreshTokenDuration = time.Hour * 24 * 7

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
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("email/password is incorrect")))
		return
	}

	if err := util.CheckPassword(loginReq.Password, account.HashedPassword); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("email/password is incorrect")))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		account.ID,
		accessTokenDuration,
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

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"account":       toAccountRes(account),
	})
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req struct {
		RefreshToken string `form:"refresh_token" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&req, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// session, err := server.store.GetSession(ctx, refreshPayload.ID)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// if session.IsBlocked {
	// 	err := fmt.Errorf("blocked session")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	// if session.Username != refreshPayload.Username {
	// 	err := fmt.Errorf("incorrect session user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	// if session.RefreshToken != req.RefreshToken {
	// 	err := fmt.Errorf("mismatched session token")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	// if time.Now().After(session.ExpiresAt) {
	// 	err := fmt.Errorf("expired session")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	accessToken, _, err := server.tokenMaker.CreateToken(
		refreshPayload.AccountID,
		accessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
