package service

import (
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

func FindDeptIdListByRoleId(roleId uint) ([]uint, error) {
	var deptIds = make([]uint, 0)
	var roleDeptList []*model.SysRoleDept
	err := database.FindListByParam(database.SearchMap{"role_id": roleId}, &roleDeptList)
	if err != nil {
		return deptIds, err
	}
	if len(roleDeptList) <= 0 {
		return deptIds, nil
	}
	for _, roleDept := range roleDeptList {
		deptIds = append(deptIds, roleDept.DeptId)
	}
	return deptIds, nil
}

func SaveOrUpdateRoleDeptBatch(roleId uint, deptIds []uint) {
	if len(deptIds) <= 0 {
		return
	}
	if roleId <= 0 {
		return
	}
	DeleteRoleDeptByRoleId(roleId)
	for _, deptId := range deptIds {
		roleDept := model.SysRoleDept{RoleId: roleId, DeptId: deptId}
		_ = database.Create(&roleDept)
	}
}

func DeleteRoleDeptByRoleId(roleId uint) {
	database.DeleteByParams(database.SearchMap{"role_id": roleId}, &model.SysRoleDept{})
}
