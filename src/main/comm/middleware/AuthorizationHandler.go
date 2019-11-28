package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/comm/config"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
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
func AuthorizationHandler(permission string) gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("<<<<<<<<<<authorization.【" + permission + "】begin>>>>>>>>>>>>>>")

		//全局環境 是否需要開啟權限攔截【必須是有注入權限攔截的接口、或全局攔截】
		if !config.CommConfig.Valid.IsOpenAuthorization {
			return
		}
		//解析 登入校驗後存入的用戶信息
		claimsData, exists := context.Get("claims")
		if !exists {
			context.JSON(http.StatusUnauthorized, utils.Error(401, "获取不到当前用户信息.[system.get.claims.not.exist]", nil))
			context.Abort()
			return
		}
		customClaims := claimsData.(*utils.CustomClaims)
		permsList, _ := service.FindPermsByUserId(customClaims.ID)
		if len(permsList) <= 0 {
			context.JSON(http.StatusUnauthorized, utils.Error(401, "用戶未授權", nil))
			context.Abort()
			return
		}
		var isPermission = false
		for _, perm := range permsList {

			if perm == permission {
				isPermission = true
				break
			}
		}
		if !isPermission {
			context.JSON(http.StatusUnauthorized, utils.Error(401, "用戶未授權", nil))
			context.Abort()
			return
		}
	}
}
