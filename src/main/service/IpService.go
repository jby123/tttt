package service

import (
	"fmt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
	"strconv"
)

/**
 * 分页获取黑白名单列表
 * @method FindIpByPage
 * @param  {[type]} name string [description]
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindIpByPage(name, order, sort string, pageNum, pageSize int) (page *database.PaginationVo) {
	searchMap := make(map[string]interface{})
	searchMap["ip"] = name
	var resultDataList []*model.SysIp
	if err := database.FindPage(searchMap, sort, sort, pageNum, pageSize).Find(&resultDataList).Error; err != nil {
		fmt.Printf("FindIpByPage.Error:%s\n", err)
		return database.Pagination(resultDataList, pageNum, pageSize, 0)
	}
	total := database.Count(searchMap, &model.SysIp{})
	return database.Pagination(resultDataList, pageNum, pageSize, total)
}

func FindIpList(ipType int) (sysIp []*model.SysIp) {
	if err := database.GetDB().Where("type = ?", ""+strconv.Itoa(ipType)+"").Find(&sysIp).Error; err != nil {
		fmt.Printf("FindIpList.Error:%s", err)
	}
	return
}

/**
 * 通过 id 获取明细
 * @method GetIpById
 * @param  {[type]}       sysIp * model.SysIp [description]
 */
func GetIpById(id uint) *model.SysIp {
	sysIp := new(model.SysIp)
	sysIp.ID = id
	if err := database.GetDB().First(sysIp).Error; err != nil {
		fmt.Printf("GetIpById.Err:%s", err)
	}

	return sysIp
}

/**
 * 通过 ip 获取明细
 * @method GetByIp
 * @param  {[type]}       sysIp * model.SysIp [description]
 */
func GetByIp(ip string) *model.SysIp {
	sysIp := new(model.SysIp)
	sysIp.Ip = ip
	if err := database.GetDB().First(sysIp).Error; err != nil {
		fmt.Printf("GetByIp.Err:%s", err)
	}

	return sysIp
}

/**
 * 通过 id 删除
 * @method DeleteIpById
 */
func DeleteIpById(id uint) {
	sysIp := new(model.SysIp)
	sysIp.ID = id
	if err := database.GetDB().Delete(sysIp).Error; err != nil {
		fmt.Printf("DeleteIpById.Err:%s", err)
	}
}

/**
 * 创建
 * @method CreateIp
 * @param  {[type]}       sysIp * model.SysIp [description]
 */
func CreateIp(sysIp *model.SysIp) *model.SysIp {

	if err := database.GetDB().Create(sysIp).Error; err != nil {
		fmt.Printf("CreateIp.Err:%s", err)
	}

	return sysIp
}

/**
 * 更新
 * @method UpdateUser
 * @param  {[type]}       sysIp * model.SysIp [description]
 */
func UpdateIp(sysIp *model.SysIp) *model.SysIp {
	if err := database.GetDB().Updates(sysIp).Error; err != nil {
		fmt.Printf("UpdateIp.Err:%s", err)
	}
	return sysIp
}
