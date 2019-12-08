package model

import (
	"time"
)

type SysUser struct {
	ID           uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt    time.Time `gorm:"column:create_time" json:"createTime" "not null VARCHAR(191)"`
	UpdatedAt    time.Time `gorm:"column:update_time" json:"updateTime" "not null VARCHAR(191)"`
	Name         string    `gorm:"column:name" json:"name" "not null VARCHAR(191)"`
	Username     string    `gorm:"column:username" json:"username" "not null VARCHAR(191)"`
	NickName     string    `gorm:"column:nick_name" json:"nickName"`
	Password     string    `gorm:"column:password" json:"password" "not null VARCHAR(191)"` //md5加密
	Phone        string    `gorm:"column:phone" json:"phone"`
	Email        string    `gorm:"column:email" json:"email"`
	HeadImg      string    `gorm:"column:head_img" json:"headImg"`
	DepartmentId uint      `gorm:"column:department_id" json:"departmentId"`
	Status       uint      `gorm:"column:status" json:"status"`
	Remark       string    `gorm:"column:remark" json:"remark"`
}

type SysUserVo struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"createTime"`
	UpdatedAt      time.Time `json:"updateTime"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	NickName       string    `json:"nickName"`
	Password       string    `json:"password"` //md5加密
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	HeadImg        string    `json:"headImg"`
	DepartmentId   uint      `json:"departmentId"`
	Status         uint      `json:"status"`
	Remark         string    `json:"remark"`
	RoleName       string    `json:"roleName"`
	DepartmentName string    `json:"departmentName"`
	RoleIdList     []uint    `json:"roleIdList"`
}

const SysUserTableName = "sys_user"

type UserJson struct {
	Username string `json:"username" validate:"required,gte=2,lte=50"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required,gte=2,lte=50"`
	RoleID   uint   `json:"roleId" validate:"required"`
}
