package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"goAdmin/src/main/utils"
)

/**
 * 授权登入 <p>https://www.bandari.net/blog/23</p>
 * 判断 token 是否有效
 * 如果有效 就获取信息并且保存到请求里面
 * @method AuthTokenHandler
 * @param  {[type]}  context       iris.Context    [description]
 */

func AuthTokenHandler() gin.HandlerFunc {
	return ValidateToken()
}

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("<<<<<<<<<<authentication.begin>>>>>>>>>>>>>>")
		token := context.Request.Header.Get("authorization")
		if token == "" {
			utils.ResultError(context, utils.BUSINESS_ERROR, "参数异常[请求未携带token，无权限访问]", nil)
			context.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				utils.ResultError(context, 401, "token.expire", nil)
				context.Abort()
				return
			}
			utils.ResultError(context, utils.SYSTEM_ERROR_CODE, err.Error(), nil)
			context.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		context.Set("claims", claims)
		fmt.Println("<<<<<<<<<<authentication.end>>>>>>>>>>>>>>")
	}
}
