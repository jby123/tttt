package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取-角色-列表
 * @method FindRoleByPage
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindRoleByPage(name, order, sort string, pageNum, pageSize int) (page *database.PaginationVo) {
	searchMap := make(map[string]interface{})
	searchMap["name"] = name
	var resultDataList []*model.SysRole

	if err := database.FindPage(searchMap, order, sort, pageNum, pageSize).Find(&resultDataList).Error; err != nil {
		fmt.Printf("FindRoleByPage.Error:%s\n", err)
		return database.Pagination(resultDataList, pageNum, pageSize, 0)
	}
	total := database.Count(searchMap, &model.SysRole{})
	return database.Pagination(resultDataList, pageNum, pageSize, total)
}

func FindRoleListByParam(searchMap map[string]interface{}) (err error, resultDataList []*model.SysRole) {
	err = database.FindListByParam(searchMap, &resultDataList)
	if err == nil {
		return nil, resultDataList
	}
	return err, resultDataList
}

func GetRoleById(id uint) (resultData *model.SysRole) {
	resultData = new(model.SysRole)
	err := database.GetById(id, resultData)
	if err != nil {
		fmt.Printf("GetRoleById.error.%s\n", err)
	}
	return resultData
}

/**
 * 创建
 * @method CreateRole
 * @param  {[type]} sysRole model.SysRole    [description]
 */
func CreateRole(sysRole *model.SysRole) error {
	err := database.Create(sysRole)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 更新
 * @method UpdateRole
 * @param  {[type]} sysRole model.SysRole    [description]
 */
func UpdateRole(sysRole *model.SysRole) error {
	err := database.UpdateById(sysRole.ID, &sysRole)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRoleById(id uint) {

	//删除角色对应的菜单
	DeleteRoleMenuByRoleId(id)
	//删除用户跟角色的关系
	DeleteUserRoleByRoleId(id)
	//删除角色跟部门关系
	DeleteRoleDeptByRoleId(id)
	//删除角色信息
	u := new(model.SysRole)
	err := database.DeleteById(id, u)
	if err != nil {
		return
	}

}
