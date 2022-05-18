package api

import "github.com/gin-gonic/gin"

func parseIdUri(ctx *gin.Context) (int32, error) {
	var idUri struct {
		Id int32 `uri:"id" binding:"required,min=1"`
	}

	err := ctx.ShouldBindUri(&idUri)

	return idUri.Id, err
}

func parsePaginateQuery(ctx *gin.Context) (int32, int32, error) {
	var paginateQuery struct {
		PageSize  int32 `form:"page_size" binding:"min=1"`
		LastID int32 `form:"last_id" binding:"min=0"`
	}

	paginateQuery.LastID = 1000_0000

	err := ctx.ShouldBindQuery(&paginateQuery)

	return paginateQuery.PageSize, paginateQuery.LastID, err
}
