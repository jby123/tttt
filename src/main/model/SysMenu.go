package model

import (
	"time"
)

type SysMenu struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	ParentId  uint      `gorm:"column:parent_id" json:"parentId"`
	Type      uint      `gorm:"column:type" json:"type"`
	Name      string    `gorm:"column:name" json:"name" "not null VARCHAR(50)"`
	Icon      string    `gorm:"column:icon" json:"icon" "not null VARCHAR(50)"`
	KeepAlive uint      `gorm:"column:keep_alive" json:"keepAlive"`
	OrderNum  uint      `gorm:"column:order_num" json:"orderNum"`
	ViewPath  string    `gorm:"column:view_path" json:"viewPath"`
	Router    string    `gorm:"column:url" json:"router"`
	Perms     string    `gorm:"column:perms" json:"perms"`
	Status    uint      `gorm:"column:status" json:"status"`
	Remark    string    `gorm:"column:remark" json:"remark"`
}
