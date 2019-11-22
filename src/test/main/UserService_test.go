package main

import (
	"fmt"
	"goAdmin/src/main/comm/config"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"testing"
)

func PreInit() {
	//初始化环境配置
	config.InitCommConfig(utils.DefaultDevelopmentRelativePath)
	//初始化mysql
	database.InitMysql(config.Config.EnvConf.ActiveConf.Active, &config.DbConf)
}
func TestGetById(t *testing.T) {
	PreInit()
	user := service.GetById(1)
	fmt.Println(user)
}
