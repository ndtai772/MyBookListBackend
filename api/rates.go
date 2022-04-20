package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)

func (server *Server) createRate(ctx *gin.Context) {
	var createRateForm struct {
		score     int32 `form:"score,omitempty"`
		createdBy int32 `form:"created_by" binding:"required"`
		bookId    int32 `form:"book_id" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&createRateForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createRateParams := db.CreateRateParams{
		BookID: createRateForm.bookId,
		CreatedBy: createRateForm.createdBy,
		RateValue: createRateForm.score,
	}

	rate, err := server.store.CreateRate(ctx, createRateParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rate)
}

func (server *Server) updateRate(ctx *gin.Context) {
	rateId, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var updateRateForm struct {
		score     int32 `form:"score,omitempty"`
	}

	if err := ctx.ShouldBindWith(&updateRateForm, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rate, err := server.store.UpdateRate(ctx, db.UpdateRateParams{
		RateValue: updateRateForm.score,
		ID: rateId,
	});

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rate)
}