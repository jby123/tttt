package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
	"goAdmin/src/main/utils"
)

/**
 * 分页获取-菜单-列表
 * @method FindMenuByPage
 * @param  {[type]} name string [description]
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindMenuByPage(name, order, sort string, pageNum, pageSize int) (page *utils.PaginationVo) {
	searchMap := make(map[string]interface{})
	searchMap["name"] = name
	var resultDataList []*model.SysMenu
	if err := database.FindPage(searchMap, sort, sort, pageNum, pageSize).Find(&resultDataList).Error; err != nil {
		fmt.Printf("FindLogByPage.Error:%s\n", err)
		return utils.Pagination(resultDataList, pageNum, pageSize, 0)
	}
	total := database.Count(searchMap, &model.SysMenu{})
	return utils.Pagination(resultDataList, pageNum, pageSize, total)
}

func FindMenuListByParam(searchMap map[string]interface{}) (err error, resultDataList []*model.SysMenu) {
	err = database.FindListByParam(searchMap, &resultDataList)
	if err == nil {
		return nil, resultDataList
	}
	return err, resultDataList
}