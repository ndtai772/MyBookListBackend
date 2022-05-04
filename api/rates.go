package api

import (
	"log"
	"net/http"

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
	// TODO: verify ownership of the rate
	rateId, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var updateRateForm struct {
		score int32 `form:"score" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&updateRateForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rate, err := server.store.UpdateRate(ctx, db.UpdateRateParams{
		RateValue: updateRateForm.score,
		ID:        rateId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	server.updateBookIndex(rate.BookID)

	ctx.JSON(http.StatusOK, rate)
}
