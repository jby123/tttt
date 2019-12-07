package service

import (
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/comm/exception"
	"goAdmin/src/main/model"
)

func SaveOrUpdateRoleDeptBatch(roleId uint, deptIds []uint) {
	if len(deptIds) <= 0 {
		return
	}

	var resultDataList []*model.SysRoleDept
	err := database.FindListByParam(database.SearchMap{"dept_id": deptIds, "role_id": roleId}, &resultDataList)
	if err != nil {
		return
	}
	if len(resultDataList) > 0 {
		return
	}
	for _, roleDept := range resultDataList {
		database.DeleteById(roleDept.ID, model.SysRoleDept{})
	}

	for _, deptId := range deptIds {
		roleDept := model.SysRoleDept{RoleId: roleId, DeptId: deptId}

		err = database.Create(roleDept)
		if err != nil {
			exception.SystemException("系统异常." + err.Error())
		}
	}

}

func DeleteRoleDeptByRoleId(roleId uint) {
	database.DeleteByParams(database.SearchMap{"role_id": roleId}, model.SysRoleDept{})
}
