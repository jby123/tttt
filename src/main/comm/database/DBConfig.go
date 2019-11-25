package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strings"
	"time"
)

var DBClient *gorm.DB

const DefaultDevelopmentEnv = "dev"

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
	if len(activeEnv) != 0 && activeEnv == DefaultDevelopmentEnv {
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

type SearchMap map[string]interface{}

/**
 * 获取列表
 * @method FindPage
 * @param  {[type]} searchMap map[string] interface{}    [description]
* @param  {[type]} searchValue string    [description]
* @param  {[type]} order string    [description]
 * @param  {[type]} orderBy string    [description]
 * @param  {[type]} relation string    [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
*/
func FindPage(searchMap map[string]interface{}, order, sort string, offset, limit int) *gorm.DB {

	if len(sort) <= 0 {
		sort = "desc"
	}
	if len(order) <= 0 {
		order = "created_at"
	}
	DBClient = DBClient.Order(" " + order + " " + sort + " ")

	searchSql, searchArgs := ParseSearchMap(searchMap)
	if len(searchSql) > 0 {
		DBClient = DBClient.Where(searchSql, searchArgs)
	}
	if offset > 0 {
		DBClient = DBClient.Offset((offset - 1) * limit)
	}
	if limit > 0 {
		DBClient = DBClient.Limit(limit)
	}
	return DBClient
}
func Count(searchMap map[string]interface{}, model interface{}) int {
	DBClient.Model(&model)
	searchSql, searchArgs := ParseSearchMap(searchMap)
	if len(searchSql) > 0 {
		DBClient = DBClient.Where(searchSql, searchArgs)
	}
	var count int = 0
	DBClient.Count(&count)
	return count
}
func ParseSearchMap(searchMap map[string]interface{}) (string, []interface{}) {
	var querySql string
	var index int = 0
	var queryArgs []interface{}
	for searchKey, searchValue := range searchMap {
		fmt.Println("searchValue", searchValue)
		if len(searchKey) > 0 {
			switch vv := searchValue.(type) {
			case string:
				if len(querySql) == 0 {
					querySql += " " + searchKey + " LIKE  ? "
				} else {
					querySql += " AND " + searchKey + " LIKE  ? "
				}
				queryArgs[index] = searchValue
			case float64:
				if len(querySql) == 0 {
					querySql += " " + searchKey + " =  ? "
				} else {
					querySql += " AND " + searchKey + " =  ? "
				}
				queryArgs[index] = searchValue
			case int:
				if len(querySql) == 0 {
					querySql += " " + searchKey + " =  ? "
				} else {
					querySql += " AND " + searchKey + " =  ? "
				}
				queryArgs[index] = searchValue
			case uint:
				if len(querySql) == 0 {
					querySql += " " + searchKey + " =  ? "
				} else {
					querySql += " AND " + searchKey + " =  ? "
				}
				queryArgs[index] = searchValue
			case []interface{}:
				var stringSlice []string
				for _, u := range vv {
					stringSlice = append(stringSlice, u.(string))
				}
				values := strings.Join(stringSlice, "_")
				if len(values) > 0 {
					if len(querySql) == 0 {
						querySql += " " + searchKey + " in (?) "
					} else {
						querySql += " AND " + searchKey + " in (?) "
					}
					queryArgs[index] = values
				}
			case nil:
				if len(querySql) == 0 {
					querySql += " " + searchKey + " IS nil "
				} else {
					querySql += " AND IS nil "
				}
			case map[string]interface{}:
				//todo 后期完成复杂实现
			default:

			}

		}
		index++
	}
	return querySql, queryArgs

}

// 关闭DB
func CloseDB() {
	fmt.Println("<<<<<<<<<<<<<DB.close.....>>>>>>")
	if DBClient != nil {
		DBClient.Close()
	}
}
