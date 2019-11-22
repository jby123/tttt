package model

import "time"

type SysLog struct {
	ID        uint      `gorm:"primary_key" sql:"auto_increment;primary_key;unique" json:"id"`
	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
	Ip        string    `gorm:"column:name" json:"name" "not null VARCHAR(50)"`
	UserId    string    `gorm:"column:user_id" json:"userId"`
	Operator  string    `gorm:"column:operator" json:"name"`
	OptUrl    string    `gorm:"column:opt_url" json:"action"`
	Params    string    `gorm:"column:password" json:"password"`
}
