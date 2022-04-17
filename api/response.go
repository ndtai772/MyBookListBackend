package api

import "github.com/gin-gonic/gin"

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func paginatedResponse(data []interface{}, total uint32, index int32) gin.H {
	return gin.H{
		"total": total,
		"data": data,
		"index": index,
	}
}