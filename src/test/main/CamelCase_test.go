package main

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
	"log"
	"reflect"
	"testing"
)

func TestCamelCase(t *testing.T) {
	// 1. 普通使用
	log.Println(database.CamelCase("AAAA"))
	log.Println(database.CamelCase("IconUrl"))
	log.Println(database.CamelCase("iconUrl"))
	log.Println(database.CamelCase("parentId"))
	log.Println(database.CamelCase("a9b9Ba"))
	log.Println(database.CamelCase("_An"))
	// s输出
	//2019/03/20 16:34:25 a_a_a_a
	//2019/03/20 16:34:25 icon_url
	//2019/03/20 16:34:25 icon_url
	//2019/03/20 16:34:25 parent_id
	//2019/03/20 16:34:25 a9b9ba
	//2019/03/20 16:34:25 Xan

	// 2. 比如数据库 User 表, 仅修改Birthday Avatar列,
	// 反射遍历字段 然后修改 拼装sql语句
	user := &model.SysUser{ID: 10, Username: "track", Password: "123456"}
	rt := reflect.ValueOf(user)
	sql := "UPDATE user SET"
	for i := 0; i < rt.Elem().NumField(); i++ {
		fieldValue := rt.Elem().Field(i)
		fieldType := rt.Elem().Type().Field(i)
		fieldName := database.CamelCase(fieldType.Name)
		if fieldName == "id" {
			continue
		}
		sql += fmt.Sprintf(" %s = %s ,", fieldName, fieldValue)
	}
	sql += " " + fmt.Sprintf("WHERE id = %d", user.ID)

	// 当然 还是有问题的. 简单举个栗子
	log.Println("SQL语句组装为:", sql)
	// SQL语句组装为: UPDATE user SET user_name = track , password = 123456 , avatar = http://baidu.com , birthday = 1999-09-09 , WHERE id = 10
}
