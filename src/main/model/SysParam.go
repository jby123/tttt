package model

import (
	"time"
)

type SysParam struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	Name      string    `gorm:"column:name" json:"name"`
	Value     string    `gorm:"column:value" json:"value"`
	Status    uint      `gorm:"column:status" json:"status"`
	Remark    string    `gorm:"column:remark" json:"remark"`
}
