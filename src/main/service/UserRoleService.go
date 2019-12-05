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

func SaveOrUpdateUserRole(userRole *model.SysUserRole) {
	var resultDataList []*model.SysUserRole
	err := database.FindListByParam(database.SearchMap{"user_id": userRole.UserId, "role_id": userRole.RoleId}, &resultDataList)
	if err != nil {
		return
	}
	if len(resultDataList) > 0 {
		return
	}
	database.Create(userRole)

}
