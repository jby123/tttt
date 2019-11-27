package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

func FindUserRoleListById(id uint) (resultDataList []*model.SysUserRole, err error) {

	err = database.FindListByParam(database.SearchMap{"user_id": id}, resultDataList)
	if err != nil {
		fmt.Printf("FindUserRoleListById.error.%s\n", err)
		return nil, err
	}
	return resultDataList, nil
}
