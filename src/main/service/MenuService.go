package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取-菜单-列表
 * @method FindMenuByPage
 * @param  {[type]} name string [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func FindMenuByPage(name, order, sort string, offset, limit int) (sysMenu []*model.SysMenu) {
	if err := database.FindPage(database.SearchMap{"name": name}, sort, sort, offset, limit).Find(&sysMenu).Error; err != nil {
		fmt.Printf("FindMenuByPage.Error:%s\n", err)
	}
	return
}
