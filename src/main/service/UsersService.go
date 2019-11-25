package service

import (
	"fmt"
	"github.com/jameskeane/bcrypt"
	"goAdmin/src/main/comm/database"
	"goAdmin/src/main/model"
)

/**
 * 分页获取用户列表
 * @method FindByPage
 * @param  {[type]} name string [description] 用户模糊查询
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func FindUserByPage(departmentId int, name, order, sort string, offset, limit int) (users []*model.SysUser) {
	searchMap := database.SearchMap{"username": name, "department_id": departmentId}
	if err := database.FindPage(searchMap, order, sort, offset, limit).Find(&users).Error; err != nil {
		fmt.Printf("FindByPage.Error:%s", err)
	}
	return
}

/**
 * 通过 id 获取 user 记录
 * @method GetById
 * @param  {[type]}       user  *SysUser [description]
 */
func GetUserById(id uint) *model.SysUser {
	user := new(model.SysUser)
	user.ID = id
	if err := database.GetDB().First(user).Error; err != nil {
		fmt.Printf("GetById.Err:%s", err)
	}
	return user
}

/**
 * 通过 username 获取 user 记录
 * @method GetByUserName
 * @param  {[type]}       user  *SysUser [description]
 */
func GetUserByUserName(username string) *model.SysUser {
	user := &model.SysUser{Username: username}
	if err := database.GetDB().First(user).Error; err != nil {
		fmt.Printf("GetByUserName.Err:%s", err)
	}

	return user
}

/**
 * 通过 id 删除用户
 * @method DeleteById
 */
func DeleteUserById(id uint) {
	u := new(model.SysUser)
	u.ID = id
	if err := database.GetDB().Delete(u).Error; err != nil {
		fmt.Printf("DeleteById.Err:%s", err)
	}
}

/**
 * 创建
 * @method Create
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func CreateUser(aul *model.UserJson) (user *model.SysUser) {
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash(aul.Password, salt)

	user = new(model.SysUser)
	user.Username = aul.Username
	user.Password = hash
	user.Name = aul.Name

	if err := database.GetDB().Create(user).Error; err != nil {
		fmt.Printf("Create.Err:%s", err)
	}

	return
}

/**
 * 更新
 * @method UpdateUser
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func UpdateUser(uj *model.UserJson, id uint) *model.SysUser {
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash(uj.Password, salt)

	user := new(model.SysUser)
	user.ID = id
	uj.Password = hash

	if err := database.GetDB().Model(user).Updates(uj).Error; err != nil {
		fmt.Printf("Update.Err:%s", err)
	}

	return user
}

/**
 * 校验用户登录
 * @method checkLogin
 * @param  {[type]}       username string [description]
 */
func checkLogin(username string) model.SysUser {
	u := model.SysUser{}
	if err := database.GetDB().Where("username = ?", username).First(&u).Error; err != nil {
		fmt.Printf("checkLogin.Err:%s", err)
	}
	return u
}

/**
 * 判断用户是否登录
 * @method Login
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func Login(username, password string) (status bool, user model.SysUser, err error) {
	user = checkLogin(username)
	if user.ID == 0 {
		return false, user, err
	} else {
		return true, user, err

	}
}
