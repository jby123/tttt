package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goAdmin/src/main/utils"
	"os"
	"time"
)

var DBClient *gorm.DB

//数据库常量
type DbConfig struct {
	Db MysqlConf `yaml:"DB"`
}

type MysqlConf struct {
	DriverName string `yaml:"DriverName"`
	Url        string `yaml:"Url"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	Dialect    string `yaml:"Dialect"`
	MaxIdle    int    `yaml:"maxIdle"`
	MaxOpen    int    `yaml:"maxOpen"`
}

// 初始化mysql
func InitMysql(activeEnv string, dbConfig *DbConfig) (err error) {
	DBClient, err = gorm.Open(dbConfig.Db.Dialect, dbConfig.Db.Url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	if len(activeEnv) != 0 && activeEnv == utils.DefaultDevelopmentEnv {
		DBClient.LogMode(true)
	}
	DBClient.DB().SetMaxIdleConns(dbConfig.Db.MaxIdle)
	DBClient.DB().SetMaxOpenConns(dbConfig.Db.MaxOpen)
	DBClient.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	err = DBClient.DB().Ping()
	/**
	 *禁用表名复数>
	 *!!!如不禁用则会出现表 y结尾边ies的问题
	 *!!!如果只是部分表需要使用源表名，请在实体类中声明TableName的构造函数
	 *
	 *func (实体名) TableName() string {
	 *   return "数据库表名"
	 *}
	 */
	DBClient.SingularTable(true)
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
func FindPage(searchKey, searchValue, orderBy string, offset, limit int) *gorm.DB {

	if len(orderBy) > 0 {
		DBClient = DBClient.Order(orderBy + "desc")
	} else {
		DBClient = DBClient.Order("created_at desc")
	}

	if len(searchKey) > 0 {
		if len(searchValue) > 0 {
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
