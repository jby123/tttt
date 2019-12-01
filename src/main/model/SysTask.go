package model

import "time"

type SysTask struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	Name      string    `gorm:"column:name" json:"name"`
	Cron      string    `gorm:"column:cron" json:"cron"`     //任务时间
	Ip        string    `gorm:"column:ip" json:"ip"`         //内网IP
	Env       string    `gorm:"column:env" json:"env"`       //环境
	Status    uint      `gorm:"column:status" json:"status"` //0 不可用、1可用
	Remark    string    `gorm:"column:remark" json:"remark"`
}
