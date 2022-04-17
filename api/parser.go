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
		Limit  int32 `query:"limit" binding:"min=1"`
		Offset int32 `query:"limit" binding:"min=0"`
	}

	err := ctx.ShouldBindQuery(&paginateQuery)

	return paginateQuery.Limit, paginateQuery.Offset, err
}
