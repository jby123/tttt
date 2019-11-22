package controller

import (
	"github.com/gin-gonic/gin"
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

	}
}

/**
 * 修改当前用户信息
 */
func UpdateCurrentUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
