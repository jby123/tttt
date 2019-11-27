package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

func FindRoleMenuListByRoleIds(roleIds []uint) (resultDataList []*model.SysRoleMenu, err error) {
	if len(roleIds) == 0 {
		return
	}
	err = database.FindListByParam(database.SearchMap{"role_id": roleIds}, resultDataList)
	if err != nil {
		fmt.Printf("FindUserRoleListById.error.%s\n", err)
		return nil, err
	}
	return resultDataList, nil
}
