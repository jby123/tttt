package model

import (
	"time"
)

type SysUser struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime" "not null VARCHAR(191)"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime" "not null VARCHAR(191)"`
	Name      string    `gorm:"column:name" json:"name" "not null VARCHAR(191)"`
	Username  string    `gorm:"column:username" json:"username" "not null VARCHAR(191)"`
	Password  string    `gorm:"column:password" json:"password" "not null VARCHAR(191)"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	RoleID    uint      `gorm:"column:role_id" json:"roleId"`
}

type UserJson struct {
	Username string `json:"username" validate:"required,gte=2,lte=50"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required,gte=2,lte=50"`
	RoleID   uint   `json:"roleId" validate:"required"`
}
