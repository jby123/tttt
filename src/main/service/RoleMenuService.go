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
