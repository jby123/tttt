package service

import (
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取日志列表
 * @method FindLogByPage
 * @param  {[type]} name string [description]
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindLogByPage(name, order, sort string, pageNum, pageSize int) (page *database.PaginationVo) {
	searchMap := make(map[string]interface{})
	searchMap["name"] = name
	var resultDataList []*model.SysLog
	if err := database.FindPage(searchMap, order, sort, pageNum, pageSize).Find(&resultDataList).Error; err != nil {
		return database.Pagination(make([]interface{}, 0), pageNum, pageSize, 0)
	}
	total := database.Count(searchMap, &model.SysLog{})
	if total > 0 {
		return database.Pagination(resultDataList, pageNum, pageSize, total)
	}
	return database.Pagination(make([]interface{}, 0), pageNum, pageSize, 0)
}

func DeleteLogsAll() {
	database.DeleteByParams(nil, &model.SysLog{})
}
