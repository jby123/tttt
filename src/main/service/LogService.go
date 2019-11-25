package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取日志列表
 * @method FindLogByPage
 * @param  {[type]} name string [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func FindLogByPage(name, order, sort string, offset, limit int) (sysLog []*model.SysLog) {
	searchMap := database.SearchMap{"name": name}
	if err := database.FindPage(searchMap, sort, sort, offset, limit).Find(&sysLog).Error; err != nil {
		fmt.Printf("FindLogByPage.Error:%s", err)
	}
	return
}
