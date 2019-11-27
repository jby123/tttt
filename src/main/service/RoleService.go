package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
	"goAdmin/src/main/utils"
)

/**
 * 分页获取-角色-列表
 * @method FindRoleByPage
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindRoleByPage(name, order, sort string, pageNum, pageSize int) (page *utils.PaginationVo) {
	searchMap := make(map[string]interface{})
	searchMap["name"] = name
	var resultDataList []*model.SysRole
	if err := database.FindPage(searchMap, sort, sort, pageNum, pageSize).Find(&resultDataList).Error; err != nil {
		fmt.Printf("FindRoleByPage.Error:%s\n", err)
		return utils.Pagination(resultDataList, pageNum, pageSize, 0)
	}
	total := database.Count(searchMap, &model.SysRole{})
	return utils.Pagination(resultDataList, pageNum, pageSize, total)
}

func FindRoleListByParam(searchMap map[string]interface{}) (err error, resultDataList []*model.SysRole) {
	err = database.FindListByParam(searchMap, &resultDataList)
	if err == nil {
		return nil, resultDataList
	}
	return err, resultDataList
}