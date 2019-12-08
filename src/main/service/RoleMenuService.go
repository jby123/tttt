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

func FindMenuIdListByRoleId(roleId uint) ([]uint, error) {
	var menuIds = make([]uint, 0)
	var roleMenuList []*model.SysRoleMenu
	err := database.FindListByParam(database.SearchMap{"role_id": roleId}, &roleMenuList)
	if err != nil {
		fmt.Printf("FindUserRoleListById.error.%s\n", err)
		return menuIds, err
	}
	if len(roleMenuList) <= 0 {
		return menuIds, nil
	}
	for _, roleMenu := range roleMenuList {
		menuIds = append(menuIds, roleMenu.MenuId)
	}
	return menuIds, nil
}

func DeleteRoleMenuByMenuId(menuId uint) {
	database.DeleteByParams(database.SearchMap{"menu_id": menuId}, &model.SysRoleMenu{})
}
func DeleteRoleMenuByRoleId(roleId uint) {
	database.DeleteByParams(database.SearchMap{"role_id": roleId}, &model.SysRoleMenu{})
}

func SaveOrUpdateRoleMenuBatch(roleId uint, menuIds []uint) {
	if len(menuIds) <= 0 {
		return
	}
	//删除角色跟菜单的关系
	database.DeleteByParams(database.SearchMap{"role_id": roleId, "menu_id": menuIds}, &model.SysRoleMenu{})
	//保存角色跟菜单的关系
	for _, menuId := range menuIds {
		roleMenu := model.SysRoleMenu{RoleId: roleId, MenuId: menuId}
		_ = database.Create(&roleMenu)
	}
}
