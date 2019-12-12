package main

import (
	"encoding/json"
	"fmt"
	"goAdmin/src/main/comm/config"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"log"
	"testing"
)

func TestFindDeptPage(t *testing.T) {
	//初始化环境配置
	config.InitCommConfig(utils.DefaultTestRelativePath)
	//初始化mysql
	database.InitDB(config.Config.EnvConf.ActiveConf.Active, &config.DbConf)

	page := service.FindDeptByPage("", "", "", 1, 10)

	data, err := json.Marshal(page)
	if err != nil {
		log.Fatalf("Json marshaling failed：%s", err)
	}
	fmt.Printf("%s\n", data)
}
