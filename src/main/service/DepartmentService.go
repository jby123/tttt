package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取部门列表
 * @method FindDeptByPage
 * @param  {[type]} name string [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func FindDeptByPage(name, order, sort string, offset, limit int) (dept []*model.SysDepartment) {
	if err := database.FindPage(database.SearchMap{"name": name}, sort, sort, offset, limit).Find(&dept).Error; err != nil {
		fmt.Printf("FindDeptByPage.Error:%s", err)
	}
	return
}
