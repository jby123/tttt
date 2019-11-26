package model

import (
	"time"
)

type SysMenu struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	ParentId  uint      `gorm:"column:parent_id" json:"parentId"`
	Name      string    `gorm:"column:name" json:"name" "not null VARCHAR(50)"`
	Url       string    `gorm:"column:url" json:"url" "not null VARCHAR(50)"`
	Icon      string    `gorm:"column:icon" json:"icon" "not null VARCHAR(50)"`
	Status    uint      `gorm:"column:status" json:"status"`
	Remark    string    `gorm:"column:remark" json:"remark"`
}
