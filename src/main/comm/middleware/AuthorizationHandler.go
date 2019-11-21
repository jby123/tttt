package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 权限认证
 * 读取用户信息
 * 读取用户资源权限
 * 判断用户是否有资源权限
 * @method AuthPermissionHandler
 * @param  {[type]}  context       gin.Context    [description]
 */
func AuthPermissionHandler(context gin.Context) {
	context.Status(http.StatusUnauthorized)
	context.Next()
}
