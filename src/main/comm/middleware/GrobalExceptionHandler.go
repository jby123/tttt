package middleware

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/comm/exception"
	"goAdmin/src/main/utils"
	"net/http"
)

/**
 * 全局異常攔截
 * TODO 需要完善err 內部拋出的狀態碼 和錯誤消息 [錯誤消息為系統國際化配置、由這邊統一轉化]
 */
func ExceptionHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case exception.Model:
					exceptionModel := err.(exception.Model)
					//能知道 err 所屬 code ,msg
					context.JSON(http.StatusBadRequest, utils.Error(exceptionModel.Code, exceptionModel.Message, nil))
					context.Abort()
				default:
					//能知道 err 所屬 code ,msg
					context.JSON(utils.SYSTEM_ERROR_CODE, utils.Error(utils.SYSTEM_ERROR_CODE, err.(string), nil))
					context.Abort()
				}
			}
		}()
		context.Next()
	}
}