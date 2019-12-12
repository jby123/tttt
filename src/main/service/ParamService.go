package service

import (
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取-參數-列表
 * @method FindParamByPage
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindParamByPage(name, order, sort string, pageNum, pageSize int) *database.PaginationVo {
	resultDataList := &[]model.SysParam{}
	model := &model.SysParam{}
	requestPage := &database.Page{
		Model:          model,
		SearchMap:      database.SearchMap{"name": name},
		ResultDataList: resultDataList,
		Pagination:     &database.PageInfo{Page: pageNum, Size: pageSize, Order: order, Sort: sort},
	}
	return database.FindCommPage(requestPage)

}

func GetParamById(id uint) *model.SysParam {
	model := &model.SysParam{}
	_, flag := database.GetById(id, &model)
	if !flag {
		return nil
	}
	return model
}

/**
 * 创建
 * @method CreateParam
 * @param  {[type]} param model.SysParam    [description]
 */
func CreateParam(param *model.SysParam) error {
	err := database.Create(param)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 更新
 * @method UpdateParam
 * @param  {[type]} param model.SysParam    [description]
 */
func UpdateParam(param *model.SysParam) error {
	err := database.Update(param)
	if err != nil {
		return err
	}
	return nil
}

func DeleteParamById(id uint) {
	param := new(model.SysParam)
	database.DeleteById(id, param)
}
