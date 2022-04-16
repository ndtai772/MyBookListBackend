package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)


func (server *Server) createAccount(ctx *gin.Context) {
	type createAccountRequest struct {
		Username string `form:"username" binding:"required,alphanum"`
		Password string `form:"password" binding:"required,min=6"`
		Email    string `form:"email" binding:"required,email"`
	}
	
	var form createAccountRequest
	if err := ctx.ShouldBindWith(&form, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Username:    form.Username,
		Email:       form.Email,
		EncodedHash: form.Password,
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
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "unimplemented",
	})
}

func (server *Server) searchAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "unimplemented",
	})
}
