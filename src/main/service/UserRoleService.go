package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

func FindUserRoleListById(id uint) ([]*model.SysUserRole, error) {
	var resultDataList []*model.SysUserRole
	err := database.FindListByParam(database.SearchMap{"user_id": id}, &resultDataList)
	if err != nil {
		fmt.Printf("FindUserRoleListById.error.%s\n", err)
		return nil, err
	}
	return resultDataList, nil
}
