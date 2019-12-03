package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
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

/**
 * 创建
 * @method CreateParam
 * @param  {[type]} dic model.SysDic    [description]
 */
func CreateDic(dic *model.SysDic) error {
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
