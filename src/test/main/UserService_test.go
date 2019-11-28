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

func PreInit() {
	//初始化环境配置
	config.InitCommConfig(utils.DefaultTestRelativePath)
	//初始化mysql
	database.InitDB(config.Config.EnvConf.ActiveConf.Active, &config.DbConf)
}
func TestGetById(t *testing.T) {
	PreInit()
	user := service.GetUserById(1)

	fmt.Println(user)
}

func TestFindUserPage(t *testing.T) {
	PreInit()
	page := service.FindUserByPage([]int{1, 2}, "", "", "", 1, 10)

	data, err := json.Marshal(page)
	if err != nil {
		log.Fatalf("Json marshaling failed：%s", err)
	}
	fmt.Printf("%s\n", data)
}

func TestFindUserCount(t *testing.T) {
	PreInit()
	count := service.FindUserCount(0, "")
	fmt.Printf("user.find.count:%d\n", count)
}
