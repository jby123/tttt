package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/xinliangnote/go-util/uuid"
	"goAdmin/src/main/comm/exception"
	"goAdmin/src/main/controller/vo"
	"goAdmin/src/main/utils"
	"io"
	"net/http"
	"os"
	"strconv"
)

/**
 * 1.获取验证码
 */
func Captcha() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		height, error := strconv.Atoi(ctx.DefaultQuery("height", "40"))
		if error != nil {
			height = 40
		}
		width, error := strconv.Atoi(ctx.DefaultQuery("width", "120"))
		if error != nil {
			height = 120
		}
		fmt.Printf("<<<<<<<<<<<Captcha.Params.height:%d, width:%d>>>>>>>>>>>>>>\n", height, width)
		var postParameters ConfigJsonBody
		//create base64 encoding captcha
		//字符,公式,验证码配置
		var config = base64Captcha.ConfigCharacter{
			Height: int(height),
			Width:  int(width),
			//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
			Mode:               base64Captcha.CaptchaModeNumber,
			ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
			ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
			IsShowHollowLine:   false,
			IsShowNoiseDot:     false,
			IsShowNoiseText:    false,
			IsShowSlimeLine:    false,
			IsShowSineLine:     false,
			CaptchaLen:         4,
		}
		captchaId, digitCap := base64Captcha.GenerateCaptcha(postParameters.Id, config)
		base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
		fmt.Printf("captchaId =%s \nbase64Png =%s\n", captchaId, base64Png)
		ctx.JSON(http.StatusOK, utils.Success(vo.CaptchaRespVo{CaptchaId: captchaId, Data: base64Png}))
	}
}

/**
 * 3.验证
 */
func VerifyCaptcha() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		captchaId := ctx.Query("captchaId")
		value := ctx.Query("value")
		if captchaId == "" || value == "" {
			ctx.JSON(http.StatusOK, utils.Error(utils.SYSTEM_ERROR_CODE, "参数有误", nil))
			return
		}
		//比较图像验证码
		verifyResult := base64Captcha.VerifyCaptcha(captchaId, value)
		if verifyResult {
			ctx.JSON(http.StatusOK, utils.Success(nil))
		} else {
			ctx.JSON(http.StatusOK, utils.Error(utils.SYSTEM_ERROR_CODE, "验证码错误", nil))
		}
	}
}

func FileUpload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//得到上传的文件
		file, header, err := ctx.Request.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
		if err != nil {
			ctx.String(http.StatusBadRequest, "Bad request")
			return
		}
		//文件的名称
		reFileName := uuid.GenUUID() + header.Filename
		//创建文件
		out, err := os.Create("static/upload_file/" + reFileName)
		//注意此处的 static/upload_file/ 不是/static/upload_file/
		if err != nil {
			exception.SystemException("upload.file.error" + err.Error())
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			exception.SystemException("upload.file.error" + err.Error())
		}
		ctx.JSON(http.StatusOK, utils.Success("static/upload_file/"+reFileName))
	}
}

//ConfigJsonBody json request body.
type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

var configC = base64Captcha.ConfigCharacter{
	Height:             60,
	Width:              240,
	Mode:               0,
	ComplexOfNoiseText: 0,
	ComplexOfNoiseDot:  0,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         6,
}
