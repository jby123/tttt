package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取-角色-列表
 * @method FindRoleByPage
 * @param  {[type]} name string [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func FindRoleByPage(name, orderBy string, offset, limit int) (sysRole []*model.SysRole) {
	if err := database.FindPage(database.SearchMap{"name": name}, orderBy, offset, limit).Find(&sysRole).Error; err != nil {
		fmt.Printf("FindRoleByPage.Error:%s\n", err)
	}
	return
}
