package main

import (
	"github.com/betacraft/yaag"
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/comm/cache"
	"goAdmin/src/main/comm/config"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/comm/middleware"
	"goAdmin/src/main/comm/routers"
	"goAdmin/src/main/model"
	"log"
	"net/http"
	"strconv"
)

func startAdminApplication() (app *gin.Engine) {
	app = gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.CorsHandler())
	//初始化环境配置
	config.InitCommConfig()
	//初始化mysql
	database.InitMysql(&config.DbConf)
	//初始化redis
	cache.InitRedis(&config.RedisConf)

	//同步模型数据表
	//如果模型表这里没有添加模型，单元测试会报错数据表不存在。
	//因为单元测试结束，会删除数据表
	database.GetDB().AutoMigrate(
		&model.User{},
	)

	routers.RegisterRoute(app)

	//api文档配置 IMPORTANT, init the middleware.
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "goAdmin",
		DocPath:   "/index.html",
		BaseUrls: map[string]string{
			"Production": "",
			"Staging":    "",
		},
	})
	//返回app
	return app
}
func main() {
	app := startAdminApplication()
	if err := http.ListenAndServe(":"+strconv.Itoa(config.Config.ServerConf.Port), app); err != nil {
		log.Fatal(err)
	}

}



