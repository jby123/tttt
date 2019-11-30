package middleware

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/model"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"net/http"
)

/**
 *IP 白名单 中间件 [某些场合下使用白名单]
 */
func IPWhiteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO 後期優化 從緩存獲取
		ipWhiteList := service.FindIpList(model.IpWhiteType)
		flag := false
		clientIp := ctx.ClientIP()
		for _, ipInfo := range ipWhiteList {
			if ipInfo.Ip == clientIp {
				flag = true
				break
			}
		}
		if !flag {
			ctx.JSON(http.StatusUnauthorized, utils.Error(401, clientIp+",不在白名单中拒绝访问", nil))
			ctx.Abort()
			return
		}
	}
}

/**
 * IP 黑名单 中间件
 */
func IPBlackHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO 後期優化 從緩存獲取
		ipBlackList := service.FindIpList(model.IpBlackType)
		flag := false
		clientIp := ctx.ClientIP()
		for _, ipInfo := range ipBlackList {
			if ipInfo.Ip == clientIp {
				flag = true
				break
			}
		}
		if flag {
			ctx.JSON(http.StatusUnauthorized, utils.Error(401, clientIp+",在黑名单中，拒绝访问", nil))
			ctx.Abort()
			return
		}
	}
}
