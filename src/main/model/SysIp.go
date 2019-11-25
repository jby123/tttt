package model

import (
	"time"
)

type SysIp struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	Ip        string    `gorm:"column:ip" json:"ip" "not null VARCHAR(50)"`
	Type      uint      `gorm:"column:type" json:"type"` //类型 1.白名单 2黑名单
	Remark    string    `gorm:"column:remark" json:"remark"`
}

const (
	IpWhiteType = 1
	IpBlackType = 2
)
