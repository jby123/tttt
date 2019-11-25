package controller

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"goAdmin/src/main/model"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"net/http"
)

type LoginReq struct {
	CaptchaId  string `json:"captchaId" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Username   string `json:"username" binding:"required"`
	VerifyCode string `json:"verifyCode" binding:"required"`
}

type LoginRespVo struct {
	Expire int    `json:"expire"`
	Token  string `json:"token"`
}

/**
 * 登入
 */
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginReq LoginReq
		error := ctx.BindJSON(&loginReq)
		if error != nil {
			fmt.Printf("login.bind.error.%s \n", error)
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		//比较验证码
		verifyResult := base64Captcha.VerifyCaptcha(loginReq.CaptchaId, loginReq.VerifyCode)
		if !verifyResult {
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "验证码错误", nil))
			return
		}
		status, user, error := service.Login(loginReq.Username, loginReq.Password)
		if error != nil {
			fmt.Printf("\n login.error.{%s}\n", error)
			return
		}
		if status {
			GenerateToken(ctx, user)
		}

	}
}

//生成令牌
func GenerateToken(c *gin.Context, user model.SysUser) {
	jwt := utils.JWT{SigningKey: []byte(utils.SignKey)}

	var expiresAt int64 = int64(utils.DEFAULT_EXPIRE_HOURE_TIME) // 过期时间 一小时
	claims := utils.CustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    utils.Issuer, // 签名的发行者
		},
	}
	token, err := jwt.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "system.error.[create.token.error]", nil))
		return
	}
	data := LoginRespVo{
		Expire: int(expiresAt),
		Token:  token,
	}
	c.JSON(http.StatusOK, utils.Success(data))
}
