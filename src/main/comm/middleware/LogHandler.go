package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 *全局日志拦截中间件
 */
func LogHandler(context *gin.Context) {
	fmt.Println("before request....<<<<<LogHandler>>>>")
	context.Next()
}

