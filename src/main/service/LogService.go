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
func FindLogByPage(name, order, sort string, pageNum, pageSize int) *database.PaginationVo {
	resultDataList := &[]model.SysLog{}
	model := &model.SysLog{}
	requestPage := &database.Page{
		Model:          model,
		SearchMap:      database.SearchMap{"name": name},
		ResultDataList: resultDataList,
		Pagination:     &database.PageInfo{Page: pageNum, Size: pageSize, Order: order, Sort: sort},
	}
	return database.FindCommPage(requestPage)

}

func DeleteLogsAll() {
	database.DeleteByParams(nil, &model.SysLog{})
}
