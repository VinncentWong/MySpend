package util

import "github.com/gin-gonic/gin"

func SendResponse(ctx *gin.Context, httpStatus int, message string, success bool, data interface{}) {
	ctx.JSON(httpStatus, gin.H{
		"message": message,
		"success": success,
		"data":    data,
	})
}
