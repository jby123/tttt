package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DBClient *gorm.DB

//数据库常量
type DbConfig struct {
	Db MysqlConf  `yaml:"DB"`
}

type MysqlConf struct {
	DriverName string `yaml:"DriverName"`
	Url        string `yaml:"Url"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	Dialect    string `yaml:"Dialect"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
}


// 初始化mysql
func InitMysql(conf *DbConfig) (err error){
	DBClient, err = gorm.Open(conf.Db.Dialect, conf.Db.Url)

	if err == nil {
		DBClient.DB().SetMaxIdleConns(conf.Db.MaxIdle)
		DBClient.DB().SetMaxOpenConns(conf.Db.MaxOpen)
		DBClient.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
		err = DBClient.DB().Ping()
	}
	return
}

// 获取mysql连接
func GetDB() *gorm.DB {
	return DBClient
}

/**
 * 获取列表
 * @method FindPage
 * @param  {[type]} searchKey string    [description]
* @param  {[type]} searchValue string    [description]
 * @param  {[type]} orderBy string    [description]
 * @param  {[type]} relation string    [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func FindPage(searchKey,searchValue, orderBy string, offset, limit int) *gorm.DB {

	if len(orderBy) > 0 {
		DBClient = DBClient.Order(orderBy + "desc")
	} else {
		DBClient = DBClient.Order("created_at desc")
	}

	if len(searchKey) > 0 {
		if len(searchValue) >0 {
			DBClient = DBClient.Where(""+searchKey+" LIKE  ?", "%"+searchValue+"%")
		}
	}

	if offset > 0 {
		DBClient = DBClient.Offset((offset - 1) * limit)
	}

	if limit > 0 {
		DBClient = DBClient.Limit(limit)
	}
	return DBClient
}


// 关闭DB
func CloseDB() {
	fmt.Println("<<<<<<<<<<<<<DB.close.....>>>>>>")
	if DBClient != nil {
		DBClient.Close()
	}
}