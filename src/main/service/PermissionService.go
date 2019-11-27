package service

import (
	"goAdmin/src/main/comm/database"
)

/**
 * 获取用户对应的菜单权限
 */
func FindPermsByUserId(userId uint) ([]string, error) {
	//获取用户对应的角色列表
	userRoleList, err := FindUserRoleListById(userId)
	if len(userRoleList) <= 0 {
		return nil, err
	}
	//获取角色列表对应的菜单权限列表
	var roleIds []uint = make([]uint, len(userRoleList))
	roleMenuList, err := FindRoleMenuListByRoleIds(roleIds)
	if len(roleMenuList) == 0 {
		return nil, err
	}
	var menuIds []uint = make([]uint, len(roleMenuList))
	err, menuList := FindMenuListByParam(database.SearchMap{"id": menuIds})
	if len(menuList) == 0 {
		return nil, err
	}
	//过滤出对应的菜单perms字段 返回
	var permsList []string
	for _, menu := range menuList {
		if len(menu.Perms) == 0 {
			continue
		}
		permsList = append(permsList, menu.Perms)
	}

	return permsList, nil
}
