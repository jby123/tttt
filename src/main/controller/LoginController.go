package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"goAdmin/src/main/utils"
	"net/http"
)

type LoginReq struct {
	CaptchaId  string `json:"captchaId" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Username   string `json:"username" binding:"required"`
	VerifyCode string `json:"verifyCode" binding:"required"`
}

/**
 * 登入
 */
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginReq LoginReq
		if ctx.BindJSON(&loginReq) != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.SYSTEM_BUSSINESS_CODE, "数据参数有误", nil))
			return
		}
		//比较验证码
		verifyResult := base64Captcha.VerifyCaptcha(loginReq.CaptchaId, loginReq.VerifyCode)
		if !verifyResult {
			ctx.JSON(http.StatusOK, utils.Error(utils.SYSTEM_ERROR_CODE, "验证码错误", nil))
			return
		}
		fmt.Printf("<<<<<<<<<<<loginReq.Params:%s>>>>>>>>>>>>>>", loginReq)
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}
