package routers

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/comm/config"
	"goAdmin/src/main/comm/middleware"
	"goAdmin/src/main/controller"
	"goAdmin/src/main/utils"
)

func RegisterRoute(application *gin.Engine) {
	//拦截器：全局日志中间件
	application.Use(gin.Logger())
	//拦截器：全局黑名单中间件
	application.Use(middleware.IPBlackHandler())
	//拦截器：全局I18N中间件
	application.Use(middleware.I18nHandler())
	//拦截器：全局Cors中间件
	application.Use(middleware.CorsHandler())
	//拦截器：全局日志中间件
	application.Use(middleware.LogHandler())
	//拦截器：全局异常中间件
	application.Use(middleware.ExceptionHandler())
	// register 404 NotFound
	application.NoRoute(middleware.GlobalNoRouteHandler)

	baseRelativePath := config.CommConfig.Iris.BasePath
	if len(baseRelativePath) == 0 {
		baseRelativePath = "/"
	}
	//1 路由组 api/admin
	route := application.Group(baseRelativePath)

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

	//1-2 路由组 api/admin/sys 下系统业务块  [需要-登入拦截-校验token][需要权限校验拦截[每個方法需要則加入權限校驗組件]]
	sys := route.Group("/sys")
	sys.Use(middleware.AuthTokenHandler())
	//日志api
	{
		log := sys.Group("/log")
		log.GET("/page", middleware.AuthorizationHandler(utils.PermsLogPage), controller.PageLogs())
	}
	//用户api
	{
		user := sys.Group("/user")
		user.GET("/page", middleware.AuthorizationHandler(utils.PermsUserPage), controller.PageUsers())
		user.GET("/list", middleware.AuthorizationHandler(utils.PermsUserList), controller.ListUsers())
		user.GET("/info", middleware.AuthorizationHandler(utils.PermsUserInfo), controller.GetUser())
		user.POST("/add", middleware.AuthorizationHandler(utils.PermsUserCreate), controller.CreateUser())
		user.POST("/update", middleware.AuthorizationHandler(utils.PermsUserUpdate), controller.UpdateUser())
		user.POST("/delete", middleware.AuthorizationHandler(utils.PermsUserDelete), controller.DeleteUser())
	}
	//角色 api
	{
		role := sys.Group("/role")
		role.GET("/page", middleware.AuthorizationHandler(utils.PermsRolePage), controller.PageRoles())
		role.GET("/list", middleware.AuthorizationHandler(utils.PermsRoleList), controller.ListRoles())
		role.GET("/info", middleware.AuthorizationHandler(utils.PermsRoleInfo), controller.GetRole())
		role.POST("/add", middleware.AuthorizationHandler(utils.PermsRoleCreate), controller.CreateRole())
		role.POST("/update", middleware.AuthorizationHandler(utils.PermsRoleUpdate), controller.UpdateRole())
		role.POST("/delete", middleware.AuthorizationHandler(utils.PermsRoleDelete), controller.DeleteRole())
	}
	//菜单 api
	{
		menu := sys.Group("/menu")
		menu.GET("/page", middleware.AuthorizationHandler(utils.PermsMenuPage), controller.PageMenus())
		menu.GET("/list", middleware.AuthorizationHandler(utils.PermsMenuList), controller.ListMenus())
		menu.GET("/info", middleware.AuthorizationHandler(utils.PermsMenuInfo), controller.GetMenu())
		menu.POST("/add", middleware.AuthorizationHandler(utils.PermsMenuCreate), controller.CreateMenu())
		menu.POST("/update", middleware.AuthorizationHandler(utils.PermsMenuUpdate), controller.UpdateMenu())
		menu.POST("/delete", middleware.AuthorizationHandler(utils.PermsMenuDelete), controller.DeleteMenu())
	}
	//部门 api
	{
		department := sys.Group("/department")
		department.GET("/page", middleware.AuthorizationHandler(utils.PermsDepartmentPage), controller.PageDepartments())
		department.GET("/list", middleware.AuthorizationHandler(utils.PermsDepartmentList), controller.ListDepartments())
		department.GET("/info", middleware.AuthorizationHandler(utils.PermsDepartmentInfo), controller.GetDepartment())
		department.POST("/add", middleware.AuthorizationHandler(utils.PermsDepartmentCreate), controller.CreateDepartment())
		department.POST("/update", middleware.AuthorizationHandler(utils.PermsDepartmentUpdate), controller.UpdateDepartment())
		department.POST("/delete", middleware.AuthorizationHandler(utils.PermsDepartmentDelete), controller.DeleteDepartment())
	}
	//系統參數 api
	{
		department := sys.Group("/param")
		department.GET("/page", controller.PageParams())
		department.GET("/info", controller.GetParam())
		department.POST("/add", controller.CreateParam())
		department.POST("/update", controller.UpdateParam())
		department.POST("/delete", controller.DeleteParam())
	}

	//數據字典 api
	{
		department := sys.Group("/dic")
		department.GET("/page", controller.PageDic())
		department.GET("/info", controller.GetDic())
		department.POST("/add", controller.CreateDic())
		department.POST("/update", controller.UpdateDic())
		department.POST("/delete", controller.DeleteDic())
	}
}
