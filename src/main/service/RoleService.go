package service

import (
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
func FindRoleByPage(name, order, sort string, pageNum, pageSize int) *database.PaginationVo {
	requestPage := &database.Page{
		Model:          &model.SysRole{},
		SearchMap:      database.SearchMap{"name": name},
		ResultDataList: &[]model.SysRole{},
		Pagination:     &database.PageInfo{Page: pageNum, Size: pageSize, Order: order, Sort: sort},
	}
	return database.FindCommPage(requestPage)

}

func FindRoleListByParam(searchMap map[string]interface{}) *[]model.SysRole {
	resultDataList := &[]model.SysRole{}
	_ = database.FindListByParam(searchMap, resultDataList)
	return resultDataList
}

func GetRoleById(id uint) *model.SysRole {
	model := &model.SysRole{}
	_, flag := database.GetById(id, &model)
	if !flag {
		return nil
	}
	return model
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
	err := database.DeleteById(id, &model.SysRole{})
	if err != nil {
		return
	}

}
