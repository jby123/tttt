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
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindDeptByPage(name, order, sort string, pageNum, pageSize int) *database.PaginationVo {
	resultDataList := &[]model.SysDepartment{}
	model := &model.SysDepartment{}
	requestPage := &database.Page{
		Model:          model,
		SearchMap:      database.SearchMap{"name": name},
		ResultDataList: resultDataList,
		Pagination:     &database.PageInfo{Page: pageNum, Size: pageSize, Order: order, Sort: sort},
	}
	return database.FindCommPage(requestPage)
}

func FindDeptListByParam(searchMap map[string]interface{}) (err error, resultDataList []*model.SysDepartment) {
	err = database.FindListByParam(searchMap, &resultDataList)
	if err == nil {
		return nil, resultDataList
	}
	return err, resultDataList
}

func GetDeptById(id uint) *model.SysDepartment {
	model := &model.SysDepartment{}
	_, flag := database.GetById(id, &model)
	if !flag {
		return nil
	}
	return model
}

/**
 * 创建
 * @method CreateDept
 * @param  {[type]} sysDept model.SysDepartment    [description]
 */
func CreateDept(sysDept *model.SysDepartment) error {
	err := database.Create(sysDept)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 更新
 * @method UpdateDept
 * @param  {[type]} sysDept model.SysDepartment    [description]
 */
func UpdateDept(sysDept *model.SysDepartment) error {
	err := database.UpdateById(sysDept.ID, sysDept)
	if err != nil {
		return err
	}
	return nil
}
func DeleteDeptById(id uint) {
	u := new(model.SysDepartment)
	err := database.DeleteById(id, u)
	if err != nil {
		fmt.Printf("DeleteDeptById.error.deptId:%d, error:%s\n", id, err)
		return
	}
	searchMap := make(map[string]interface{})
	searchMap["department_id"] = id
	DeleteUserByParams(searchMap)

}
