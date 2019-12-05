package model

import (
	"time"
)

type SysRoleDept struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	RoleId    uint      `gorm:"column:role_id" json:"roleId"`
	DeptId    uint      `gorm:"column:dept_id" json:"deptId"`
	Remark    string    `gorm:"column:remark" json:"remark"`
}
