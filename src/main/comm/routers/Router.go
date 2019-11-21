package routers

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/comm/middleware"
	"goAdmin/src/main/utils"
	"net/http"
)

func RegisterRoute(application *gin.Engine) {
	//拦截器：全局日志中间件
	application.Use(gin.Logger())
	//拦截器：全局异常恢复中间件
	application.Use(gin.Recovery())
	//拦截器：全局Cors中间件
	application.Use(middleware.CorsHandler())
	//拦截器：全局日志中间件
	application.Use(middleware.LogHandler)


	// register 404 NotFound
	application.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.Error(http.StatusBadRequest, "NOT.FOUND", nil))
	})


	// users Api
	{
		route := application.Group("/users")
		route.GET("api/v1/articles", nil)

	}

}
