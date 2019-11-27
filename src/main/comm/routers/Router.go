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
	//拦截器：全局黑名单中间件
	application.Use(middleware.IPBlackHandler())
	//拦截器：全局I18N中间件
	application.Use(middleware.I18nHandler())
	//拦截器：全局Cors中间件
	application.Use(middleware.CorsHandler())
	//拦截器：全局日志中间件
	application.Use(middleware.LogHandler())

	// register 404 NotFound
	application.NoRoute(middleware.GlobalNoRouteHandler)

	//1 路由组 api/admin
	route := application.Group("/api/admin")

	//1-1 路由组 api/admin/comm 公共业务网块
	comm := route.Group("/comm")
	//公共业务块api
	{
		comm.POST("/login", controller.Login())    //登入
		comm.GET("/captcha", controller.Captcha()) //获取验证码
		//comm.GET("/verifyCaptcha", controller.VerifyCaptcha()) //验证-验证码

		//登入后 获取  [需要校验token]
		comm.GET("/permmenu", middleware.AuthTokenHandler(), controller.CurrentAuthorizationMenus())
		comm.GET("/person", middleware.AuthTokenHandler(), controller.CurrentUserInfo())               //当前用户信息
		comm.POST("/person-update", middleware.AuthTokenHandler(), controller.UpdateCurrentUserInfo()) //修改当前用户信息[修改密码、头像]
	}

	//1-2 路由组 api/admin/sys 系统业务块  [需要校验token][需要权限校验拦截、登入拦截]
	sys := route.Group("/sys")
	sys.Use(middleware.AuthTokenHandler())
	//日志api
	{
		log := sys.Group("/log")
		log.GET("/page", middleware.AuthorizationHandler("sys:log:page"), controller.PageLogs())
	}
	//用户api
	{
		user := sys.Group("/user")
		user.Use(middleware.AuthorizationHandler(""))
		user.GET("/page", controller.PageUsers())
		user.GET("/list", controller.ListUsers())
		user.GET("/info", controller.GetUser())
		user.POST("/add", controller.CreateUser())
		user.POST("/update", controller.UpdateUser())
		user.POST("/delete", controller.DeleteUser())
	}
	//角色 api
	{
		role := sys.Group("/role")
		role.Use(middleware.AuthorizationHandler(""))
		role.GET("/page", controller.PageRoles())
		role.GET("/list", controller.ListRoles())
		role.GET("/info", nil)
		role.POST("/add", nil)
		role.POST("/update", nil)
		role.POST("/delete", nil)
	}
	//菜单 api
	{
		menu := sys.Group("/menu")
		menu.Use(middleware.AuthorizationHandler(""))
		menu.GET("/page", controller.PageMenus())
		menu.GET("/list", controller.ListMenus())
		menu.GET("/info", nil)
		menu.POST("/add", nil)
		menu.POST("/update", nil)
		menu.POST("/delete", nil)
	}
	//部门 api
	{
		department := sys.Group("/department")
		department.Use(middleware.AuthorizationHandler(""))
		department.GET("/page", controller.PageDepartments())
		department.GET("/list", controller.ListDepartments())
		department.GET("/info", nil)
		department.POST("/add", nil)
		department.POST("/update", nil)
		department.POST("/delete", nil)
	}

}
