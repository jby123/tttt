package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginReq struct {
	CaptchaId  string `json:"captchaId"`
	Password   string `json:"password"`
	Username   string `json:"username"`
	VerifyCode string `json:"verifyCode"`
}

/**
 * 登入
 */
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginReq LoginReq
		ctx.BindJSON(&loginReq)
		fmt.Printf("<<<<<<<<<<<loginReq.Params:%s>>>>>>>>>>>>>>", loginReq)
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}
