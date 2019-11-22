package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 分页获取日志列表
 */
func LogPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}
