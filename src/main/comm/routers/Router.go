package routers

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/comm/middleware"
	"goAdmin/src/main/controller"
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
	application.NoRoute(middleware.GlobalNoRouteHandler)

	//1 路由组 api/admin
	route := application.Group("/api/admin")

	//1-1 路由组 api/admin/comm 公共业务网块
	comm := route.Group("/comm")
	//公共业务块api
	{
		comm.GET("/captcha", controller.Captcha()) //获取验证码
		comm.POST("/login", controller.Login())    //登入

		//登入后 获取  [需要校验token]
		comm.GET("/permmenu", controller.CurrentAuthorizationMenus())
		comm.GET("/person", controller.CurrentUserInfo())               //当前用户信息
		comm.POST("/person-update", controller.UpdateCurrentUserInfo()) //修改当前用户信息[修改密码、头像]
	}

	//1-2 路由组 api/admin/sys 系统业务块  [需要校验token][需要权限校验拦截、登入拦截]
	sys := route.Group("/sys")
	sys.Use(middleware.AuthTokenHandler())
	//日志api
	{
		log := sys.Group("/log")
		log.GET("/page", controller.LogPage())
	}
	//用户api
	{
		user := sys.Group("/user")
		user.GET("/page", controller.PageUsers())
		user.GET("/list", controller.GetAllUsers())
		user.GET("/info", controller.GetUser())
		user.POST("/add", controller.CreateUser())
		user.POST("/update", controller.UpdateUser())
		user.POST("/delete", controller.DeleteUser())
	}
	//角色 api
	{
		role := sys.Group("/role")
		role.GET("/page", nil)
		role.GET("/list", nil)
		role.GET("/info", nil)
		role.POST("/add", nil)
		role.POST("/update", nil)
		role.POST("/delete", nil)
	}
	//角色 api
	{
		menu := sys.Group("/menu")
		menu.GET("/page", nil)
		menu.GET("/list", nil)
		menu.GET("/info", nil)
		menu.POST("/add", nil)
		menu.POST("/update", nil)
		menu.POST("/delete", nil)
	}
	//部门 api
	{
		department := sys.Group("/department")
		department.GET("/page", nil)
		department.GET("/list", nil)
		department.GET("/info", nil)
		department.POST("/add", nil)
		department.POST("/update", nil)
		department.POST("/delete", nil)
	}

}
