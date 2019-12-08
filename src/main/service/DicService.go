package service

import (
	"fmt"
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
func FindDicByPage(name, order, sort string, pageNum, pageSize int) (page *database.PaginationVo) {
	searchMap := make(map[string]interface{})
	searchMap["name"] = name
	var resultDataList []*model.SysDic
	if err := database.FindPage(searchMap, sort, sort, pageNum, pageSize).Find(&resultDataList).Error; err != nil {
		fmt.Printf("FindDicByPage.Error:%s\n", err)
		return database.Pagination(resultDataList, pageNum, pageSize, 0)
	}
	total := database.Count(searchMap, &model.SysDic{})
	return database.Pagination(resultDataList, pageNum, pageSize, total)
}

func GetDicById(id uint) (resultData *model.SysDic) {
	resultData = new(model.SysDic)
	err := database.GetById(id, resultData)
	if err != nil {
		fmt.Printf("GetDicById.error.%s\n", err)
	}
	return resultData
}

func GetDicByName(name string) *model.SysDic {
	sysDic := new(model.SysDic)
	_ = database.FindListByParam(database.SearchMap{"name": database.Condition{Operate: database.Equals, Value: name}}, sysDic)
	return sysDic
}

/**
 * 创建
 * @method CreateParam
 * @param  {[type]} dic model.SysDic    [description]
 */
func CreateDic(dic *model.SysDic) error {

	sysDic := model.SysDic{}
	isExist := database.Count(database.SearchMap{"name": database.Condition{Operate: database.Equals, Value: dic.Name}}, &sysDic)
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
	param := new(model.SysDic)
	database.DeleteById(id, param)
}
