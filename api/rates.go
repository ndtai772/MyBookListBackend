package api

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/token"
)

func (server *Server) createRate(ctx *gin.Context) {
	var createRateForm struct {
		Score  int32 `form:"score" binding:"required"`
		BookID int32 `form:"book_id" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&createRateForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	log.Println(createRateForm)

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	createRateParams := db.CreateRateParams{
		BookID:    createRateForm.BookID,
		CreatedBy: authPayload.AccountID,
		RateValue: createRateForm.Score,
	}

	rate, err := server.store.CreateRate(ctx, createRateParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	server.updateBookIndex(createRateForm.BookID)

	ctx.JSON(http.StatusOK, rate)
}

func (server *Server) updateRate(ctx *gin.Context) {
	rateId, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rate, err := server.store.GetRate(ctx, rateId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if rate.CreatedBy != authPayload.AccountID {
		ctx.JSON(http.StatusForbidden, errorResponse(errors.New("you don't own this rate")))
		return
	}

	var updateRateForm struct {
		Score int32 `form:"score" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&updateRateForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rate, err = server.store.UpdateRate(ctx, db.UpdateRateParams{
		RateValue: updateRateForm.Score,
		ID:        rateId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	server.updateBookIndex(rate.BookID)

	ctx.JSON(http.StatusOK, rate)
}

func (server *Server) checkRate(ctx *gin.Context) {
	var req struct {
		AccountID int `form:"account_id" binding:"required"`
		BookID    int `form:"book_id" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&req, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.AccountID != int32(req.AccountID) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	rate, err := server.store.CheckRate(ctx, db.CheckRateParams{
		BookID:    int32(req.BookID),
		CreatedBy: int32(req.AccountID),
	})

	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows in result set") {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("you didn't rate this book before")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rate)
}
