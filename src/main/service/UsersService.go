package service

import (
	"encoding/json"
	"fmt"
	"goAdmin/src/main/comm/cache"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/comm/exception"
	"goAdmin/src/main/model"
	"goAdmin/src/main/utils"
	"strconv"
	"strings"
)

/**
 * 分页获取用户列表
 * @method FindByPage
 * @param  {[type]} departmentIds []int [description]
 * @param  {[type]} name string [description] 用户模糊查询
 * @param  {[type]} order string [description]
 * @param  {[type]} sort string [description]
 * @param  {[type]} pageNum int    [description]
 * @param  {[type]} pageSize int    [description]
 */
func FindUserByPage(departmentIds []int, name, order, sort string, pageNum, pageSize int) *database.PaginationVo {

	resultDataList := &[]model.SysUserVo{}

	//1.統計
	searchMap := make(map[string]interface{})
	searchMap["username"] = name
	searchMap["department_id"] = departmentIds
	total := database.Count(searchMap, &model.SysUser{})
	if total == 0 {
		return database.Pagination(make([]interface{}, 0), pageNum, pageSize, 0)
	}
	queryArgs := make([]interface{}, 0)
	var sql string = ` SELECT u.*, GROUP_CONCAT(r.name) AS role_name, d.name AS department_name `
	sql += ` FROM sys_user AS u `
	sql += ` LEFT JOIN sys_user_role AS ur ON u.id = ur.user_id `
	sql += ` LEFT JOIN sys_role AS r ON ur.role_id = r.id `
	sql += ` LEFT JOIN sys_department AS d  ON d.id = u.department_id `
	sql += ` WHERE 1=1 `
	if len(name) > 0 {
		sql += ` AND username like ? `
		queryArgs = append(queryArgs, "%"+name+"%")
	}
	if len(departmentIds) > 0 {
		sql += ` AND department_id IN (?) `
		queryArgs = append(queryArgs, departmentIds)
	}
	sql += ` GROUP BY u.id  `
	sql += ` {filterLimit} `

	offset := (pageNum - 1) * pageSize
	limit := "limit " + strconv.Itoa(offset) + ", " + strconv.Itoa(pageSize) + ""
	sql = strings.Replace(sql, "{filterLimit}", limit, -1)

	err := database.GetDB().Raw(sql, queryArgs...).Scan(resultDataList).Error
	if err != nil {
		fmt.Printf("FindByPage.Error:%s\n", err)
		return database.Pagination(make([]interface{}, 0), pageNum, pageSize, 0)
	}
	return database.Pagination(resultDataList, pageNum, pageSize, total)
}

func FindUserCount(departmentId int, name string) int {
	searchMap := make(map[string]interface{})
	searchMap["username"] = name
	if departmentId > 0 {
		searchMap["department_id"] = departmentId
	}
	return database.Count(searchMap, &model.SysUser{})
}

/**
 * 通过 id 获取 user 记录
 * @method GetUserById
 * @param  {[type]}       user  *SysUser [description]
 */
func GetUserById(id uint) *model.SysUser {
	model := &model.SysUser{}
	//获取缓存
	value, _ := cache.Get(fmt.Sprintf(utils.CacheUserKey, id))
	if len(value) <= 0 {
		_, flag := database.GetById(id, model)
		if !flag {
			return nil
		}
		value, _ := json.Marshal(model)
		cache.Set(fmt.Sprintf(utils.CacheUserKey, id), string(value), utils.DEFAULT_EXPIRE_DAY_TIME)
	} else {
		//添加缓存
		_ = json.Unmarshal([]byte(value), model)
	}

	return model
}

/**
 * 通过 username 获取 user 记录
 * @method GetUserByUserName
 * @param  {[type]}       user  *SysUser [description]
 */
func GetUserByUserName(username string) (*model.SysUser, bool, error) {
	model := &model.SysUser{}
	if err := database.GetDB().Where("username = ?", username).First(model).Error; err != nil {
		return nil, false, err
	}

	return model, true, nil
}

func DeleteUserByParams(searchMap map[string]interface{}) {
	//TODO
}

/**
 * 通过 id 删除用户
 * @method DeleteUserById
 */
func DeleteUserById(id uint) error {
	_ := database.DeleteById(id, &model.SysUser{})
	_, _ = cache.Delete(fmt.Sprintf(utils.CacheUserKey, id))
	return nil
}

func DeleteUserByIds(ids []int) error {
	if len(ids) <= 0 {
		return nil
	}
	_ := database.DeleteByIds(ids, &model.SysUser{})
	for id := range ids {
		_, _ = cache.Delete(fmt.Sprintf(utils.CacheUserKey, id))
	}
	return nil
}

/**
 * 创建
 * @method CreateUser
 * @param  {[type]} sysUser model.SysUser [description]
 */
func CreateUser(sysUser *model.SysUser) error {
	if len(sysUser.Password) == 0 {
		sysUser.Password = utils.DefaultPassword
	}
	sysUser.Password = utils.EncodeMD5(sysUser.Password)
	err := database.Create(sysUser)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 更新
 * @method UpdateUser
 * @param  {[type]} sysUser model.SysUser [description]
 */
func UpdateUser(sysUser *model.SysUser) error {
	if len(sysUser.Password) == 0 {
		sysUser.Password = utils.DefaultPassword
	}
	sysUser.Password = utils.EncodeMD5(sysUser.Password)

	err := database.UpdateById(sysUser.ID, sysUser)
	if err != nil {
		return err
	}
	_, _ = cache.Delete(fmt.Sprintf(utils.CacheUserKey, sysUser.ID))
	return nil
}

/**
 * 判断用户是否登录
 * @method Login
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func Login(username, password string) (bool, *model.SysUser, error) {
	user, status, err := GetUserByUserName(username)
	if !status {
		exception.LoginAccountException()
	} else {
		if !utils.VerifyMD5(password, user.Password) {
			exception.LoginAccountException()
		}
	}
	return true, user, err
}
