package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 授权登入
 * 判断 token 是否有效
 * 如果有效 就获取信息并且保存到请求里面
 * @method AuthTokenHandler
 * @param  {[type]}  context       iris.Context    [description]
 */

func AuthTokenHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("<<<<<<<<<<authentication.begin>>>>>>>>>>>>>>")
		context.Status(http.StatusOK)
		context.Next()
		fmt.Println("<<<<<<<<<<authentication.end>>>>>>>>>>>>>>")
	}
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")
		if token == "" {
			c.JSON(401, gin.H{
				"message": "Token required",
			})
			c.Abort()
			return
		}
		if token != "accesstoken" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
