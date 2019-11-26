package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"net/http"
)

func Logout(ctx gin.Context) {
	ctx.Status(http.StatusOK)
	ctx.JSON(utils.SuccessCode, nil)
}

/**
 * 获取当前用户对应的用户权限列表
 */
func CurrentAuthorizationMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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

		fmt.Println("CurrentUserInfo.claims....userId", customClaims.ID)

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
