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

func FindUserRoleByRoleId(roleId uint) []model.SysUserRole {
	var resultDataList []model.SysUserRole
	err := database.FindListByParam(database.SearchMap{"role_id": roleId}, &resultDataList)
	if err != nil {
		return make([]model.SysUserRole, 0)
	}
	if len(resultDataList) <= 0 {
		return make([]model.SysUserRole, 0)
	}
	return resultDataList
}

func FindRoleIdListByUserId(userId uint) []uint {
	var resultDataList []model.SysUserRole
	err := database.FindListByParam(database.SearchMap{"user_id": userId}, &resultDataList)
	if err != nil {
		return make([]uint, 0)
	}
	if len(resultDataList) <= 0 {
		return make([]uint, 0)
	}
	roleIds := make([]uint, 0)
	for _, userRole := range resultDataList {
		roleIds = append(roleIds, userRole.RoleId)
	}
	return roleIds
}
func DeleteUserRoleByRoleId(roleId uint) {
	database.DeleteByParams(database.SearchMap{"role_id": roleId}, &model.SysUserRole{})
}
func DeleteUserRoleByUserId(userId uint) {
	database.DeleteByParams(database.SearchMap{"user_id": userId}, &model.SysUserRole{})
}
