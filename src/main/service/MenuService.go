package service

import (
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取-菜单-列表
 * @method FindMenuByPage
 * @param  {[type]} name string [description]
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindMenuByPage(name, order, sort string, pageNum, pageSize int) *database.PaginationVo {
	resultDataList := &[]model.SysMenu{}
	model := &model.SysMenu{}
	requestPage := &database.Page{
		Model:          model,
		SearchMap:      database.SearchMap{"name": name},
		ResultDataList: resultDataList,
		Pagination:     &database.PageInfo{Page: pageNum, Size: pageSize, Order: order, Sort: sort},
	}
	return database.FindCommPage(requestPage)

}

func FindMenuListByParam(searchMap map[string]interface{}) (err error, resultDataList []*model.SysMenu) {
	resultDataList = make([]*model.SysMenu, 0)
	err = database.FindListByParam(searchMap, &resultDataList)
	if err == nil {
		return nil, resultDataList
	}
	return err, resultDataList
}

func GetMenuById(id uint) *model.SysMenu {
	model := &model.SysMenu{}
	_, flag := database.GetById(id, &model)
	if !flag {
		return nil
	}
	return model
}

/**
 * 创建
 * @method CreateMenu
 * @param  {[type]} sysMenu model.SysMenu    [description]
 */
func CreateMenu(sysMenu *model.SysMenu) error {
	err := database.Create(sysMenu)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 更新
 * @method UpdateMenu
 * @param  {[type]} sysMenu model.SysMenu    [description]
 */
func UpdateMenu(sysMenu *model.SysMenu) error {
	err := database.UpdateById(sysMenu.ID, sysMenu)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMenuById(id uint) {
	err := database.DeleteById(id, &model.SysMenu{})
	if err != nil {
		return
	}
	/**
	 * 清除角色与 菜单的关系
	 */
	DeleteRoleMenuByMenuId(id)

}
