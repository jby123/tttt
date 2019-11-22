package model

import (
	"time"
)

type SysUserDept struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	UserId    uint      `gorm:"column:user_id" json:"userId"`
	DeptId    uint      `gorm:"column:dept_id" json:"deptId"`
	Remark    string    `gorm:"column:remark" json:"remark"`
}
