package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
	"goAdmin/src/main/utils"
)

/**
 * 分页获取-參數-列表
 * @method FindParamByPage
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindParamByPage(name, order, sort string, pageNum, pageSize int) (page *utils.PaginationVo) {
	searchMap := make(map[string]interface{})
	searchMap["name"] = name
	var resultDataList []*model.SysParam
	if err := database.FindPage(searchMap, sort, sort, pageNum, pageSize).Find(&resultDataList).Error; err != nil {
		fmt.Printf("FindParamByPage.Error:%s\n", err)
		return utils.Pagination(resultDataList, pageNum, pageSize, 0)
	}
	total := database.Count(searchMap, &model.SysRole{})
	return utils.Pagination(resultDataList, pageNum, pageSize, total)
}

func GetParamById(id uint) (resultData *model.SysParam) {
	resultData = new(model.SysParam)
	err := database.GetById(id, resultData)
	if err != nil {
		fmt.Printf("GetParamById.error.%s\n", err)
	}
	return resultData
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
