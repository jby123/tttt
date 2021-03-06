package model

import (
	"time"
)

type SysDepartment struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	Name      string    `gorm:"column:name" json:"name" "not null VARCHAR(50)"`
	ParentId  uint      `gorm:"column:parent_id" json:"parentId"`
	Status    uint      `gorm:"column:status" json:"status"`
	OrderNum  uint      `gorm:"column:order_num" json:"orderNum"`
	Remark    string    `gorm:"column:remark" json:"remark"`
}
