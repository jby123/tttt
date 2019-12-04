package controller

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"net/http"
)

func Logout(ctx gin.Context) {
	ctx.JSON(http.StatusOK, utils.Success(nil))
}

/**
 * 获取当前用户对应的用户权限列表
 */
func CurrentAuthorizationMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimsData, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "获取不到当前用户信息.[system.get.claims.not.exist]", nil))
			return
		}
		customClaims := claimsData.(*utils.CustomClaims)

		//1.获取所有的菜单列表
		_, menuList := service.FindMenuListByParam(nil)

		//2.获取用户对应的权限列表
		permsList, _ := service.FindPermsByUserId(customClaims.ID)
		var resultMap map[string]interface{} = make(map[string]interface{})
		resultMap["menus"] = menuList
		resultMap["perms"] = permsList
		ctx.JSON(http.StatusOK, utils.Success(resultMap))
	}
}

/**
 * 获取当前用户信息
 */
func CurrentUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimsData, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "获取不到当前用户信息.[system.get.claims.not.exist]", nil))
			return
		}
		customClaims := claimsData.(*utils.CustomClaims)
		ctx.JSON(http.StatusOK, utils.Success(service.GetUserById(customClaims.ID)))
	}
}

/**
 * 修改当前用户信息
 */
func UpdateCurrentUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
