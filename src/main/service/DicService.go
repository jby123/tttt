package service

import (
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/comm/exception"
	"goAdmin/src/main/model"
	"goAdmin/src/main/utils"
	"strings"
)

/**
 * 分页获取-字典-列表
 * @method FindDicByPage
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindDicByPage(name, order, sort string, pageNum, pageSize int) *database.PaginationVo {
	resultDataList := &[]model.SysDic{}
	model := &model.SysDic{}
	requestPage := &database.Page{
		Model:          model,
		SearchMap:      database.SearchMap{"name": name},
		ResultDataList: resultDataList,
		Pagination:     &database.PageInfo{Page: pageNum, Size: pageSize, Order: order, Sort: sort},
	}
	return database.FindCommPage(requestPage)

}

func GetDicById(id uint) *model.SysDic {
	model := &model.SysDic{}
	_, flag := database.GetById(id, &model)
	if !flag {
		return nil
	}
	return model
}

func GetDicByName(name string) *model.SysDic {
	model := &model.SysDic{}
	_ = database.FindListByParam(database.SearchMap{"name": database.Condition{Operate: database.Equals, Value: name}}, model)
	return model
}

/**
 * 创建
 * @method CreateParam
 * @param  {[type]} dic model.SysDic    [description]
 */
func CreateDic(dic *model.SysDic) error {

	isExist := database.Count(database.SearchMap{"name": database.Condition{Operate: database.Equals, Value: dic.Name}}, &model.SysDic{})
	if isExist > 0 {
		exception.BusinessException(utils.BUSINESS_ERROR, "dic.exist")
	}
	err := database.Create(dic)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 更新
 * @method UpdateDic
 * @param  {[type]} dic model.SysDic    [description]
 */
func UpdateDic(dic *model.SysDic) error {

	dbDic := GetDicById(dic.ID)
	isExit := dbDic != nil && !strings.EqualFold(dbDic.Name, dic.Name)
	if isExit {
		isExist := database.Count(database.SearchMap{"name": database.Condition{Operate: database.Equals, Value: dic.Name}}, &model.SysDic{})
		if isExist > 0 {
			exception.BusinessException(utils.BUSINESS_ERROR, "dic.exist")
		}
	}
	err := database.Update(dic)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDicById(id uint) {
	database.DeleteById(id, &model.SysDic{})
}
