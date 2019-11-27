package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/comm/cache"
	"goAdmin/src/main/comm/config"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/comm/routers"
	"goAdmin/src/main/model"
	"goAdmin/src/main/utils"
	"log"
	"net/http"
	"os"
	"strconv"
)

func startAdminApplication() (app *gin.Engine) {
	app = gin.New()
	//初始化环境配置
	config.InitCommConfig(utils.DefaultDevelopmentRelativePath)
	//初始化mysql
	database.InitDB(config.Config.EnvConf.ActiveConf.Active, &config.DbConf)
	//初始化redis
	cache.InitRedis(&config.RedisConf)

	//同步模型数据表
	//如果模型表这里没有添加模型，单元测试会报错数据表不存在。
	//因为单元测试结束，会删除数据表
	database.GetDB().AutoMigrate(
		&model.SysUser{},
		&model.SysLog{},
		&model.SysIp{},
		&model.SysDepartment{},
		&model.SysRole{},
		&model.SysMenu{},
		&model.SysRoleMenu{},
		&model.SysUserRole{},
	)
	routers.RegisterRoute(app)
	//返回app
	return app
}
func main() {
	app := startAdminApplication()
	if err := http.ListenAndServe(":"+strconv.Itoa(config.Config.ServerConf.Port), app); err != nil {
		log.Fatal(err)
	}
	fmt.Println("|-----------------------------------|")
	fmt.Println("|            go-admin             |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Admin Server Start Successful  |")
	fmt.Println("|    Port" + strconv.Itoa(config.Config.ServerConf.Port) + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")
}
