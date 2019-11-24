package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 *全局日志拦截中间件
 */
func LogHandler(context *gin.Context) {
	fmt.Println("before request....<<<<<LogHandler.begin>>>>")
	context.Status(http.StatusOK)
	context.Next()
	fmt.Println("after request ....<<<<LogHandler.end>>>>>")
}
