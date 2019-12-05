package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

func FindRoleMenuListByRoleIds(roleIds []uint) ([]*model.SysRoleMenu, error) {

	var resultDataList []*model.SysRoleMenu
	err := database.FindListByParam(database.SearchMap{"role_id": roleIds}, &resultDataList)
	if err != nil {
		fmt.Printf("FindUserRoleListById.error.%s\n", err)
		return nil, err
	}
	return resultDataList, nil
}

func DeleteRoleMenuByParams(searchMap map[string]interface{}) {
	//TODO
}

func SaveOrUpdateRoleMenuBatch(roleId uint, menuIds []uint) {
	if len(menuIds) <= 0 {
		return
	}
	//删除角色跟菜单的关系
	database.DeleteByParams(database.SearchMap{"role_id": roleId, "menu_id": menuIds}, model.SysRoleMenu{})
	//保存角色跟菜单的关系
	for _, menuId := range menuIds {
		roleMenu := model.SysRoleMenu{RoleId: roleId, MenuId: menuId}
		_ = database.Create(roleMenu)
	}
}
