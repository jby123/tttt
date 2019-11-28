package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"os"
	"strings"
	"time"
)

var dbClient *gorm.DB

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

// 初始化 DB 连接
func InitDB(activeEnv string, dbConfig *DbConfig) (err error) {
	dbClient, err = gorm.Open(dbConfig.Db.Dialect, dbConfig.Db.Url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	if len(activeEnv) != 0 && activeEnv == DefaultDevelopmentEnv {
		dbClient.LogMode(true)
	}
	dbClient.DB().SetMaxIdleConns(dbConfig.Db.MaxIdle)
	dbClient.DB().SetMaxOpenConns(dbConfig.Db.MaxOpen)
	dbClient.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute) //心跳時間設置為MySQL的一半
	//注册 db回调钩子操作
	dbClient.Callback().Create().Replace("gorm:create_time", updateTimeStampForCreateCallback)
	dbClient.Callback().Update().Replace("gorm:update_time", updateTimeStampForUpdateCallback)
	err = dbClient.DB().Ping()
	/**
	 *禁用表名复数>
	 *!!!如不禁用则会出现表 y结尾边ies的问题
	 *!!!如果只是部分表需要使用源表名，请在实体类中声明TableName的构造函数
	 *
	 *func (实体名) TableName() string {
	 *   return "数据库表名"
	 *}
	 */
	dbClient.SingularTable(true)
	return
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}

// 获取mysql连接
func GetDB() *gorm.DB {
	return dbClient
}

type SearchMap map[string]interface{}

/**
 * 获取列表
 * @method FindPage
 * @param  {[type]} searchMap map[string] interface{}    [description]
* @param  {[type]} searchValue string    [description]
* @param  {[type]} order string    [description]
 * @param  {[type]} sort string    [description]
 * @param  {[type]} relation string    [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
*/
func FindPage(searchMap map[string]interface{}, order, sort string, offset, limit int) *gorm.DB {

	if len(sort) <= 0 {
		sort = "desc"
	}
	if len(order) <= 0 {
		order = "create_time"
	}
	dbClient.Order(" " + order + " " + sort + " ")

	searchSql, searchArgs := ParseSearchMap(searchMap)
	if len(searchSql) > 0 {
		dbClient.Where(searchSql, searchArgs)
	}
	if offset > 0 {
		dbClient.Offset((offset - 1) * limit)
	}
	if limit > 0 {
		dbClient.Limit(limit)
	}
	return dbClient
}

func FindListByParam(searchMap map[string]interface{}, resultDataList interface{}) error {
	searchSql, searchArgs := ParseSearchMap(searchMap)
	if len(searchSql) > 0 {
		dbClient.Where(searchSql, searchArgs)
	}
	if err := dbClient.Find(resultDataList).Error; err != nil {
		fmt.Printf("FindListByParam.Err:%s\n", err)
		return err
	}
	return nil
}

func Count(searchMap map[string]interface{}, model interface{}) int {
	dbClient.Model(model)
	searchSql, searchArgs := ParseSearchMap(searchMap)
	if len(searchSql) > 0 {
		dbClient.Where(searchSql, searchArgs)
	}
	var total int = 0
	dbClient.Count(&total)
	return total
}

/**
 * 通过 id 获取记录
 * @method GetById
 * @param  {[type]} model interface{} [description]
 */
func GetById(id uint, model interface{}) error {
	if err := dbClient.Where("id = ?", id).First(model).Error; err != nil {
		fmt.Printf("GetById.Err:%s\n", err)
		return err
	}
	return nil
}

func DeleteById(id uint, model interface{}) error {
	if err := dbClient.Where("id = ?", id).Delete(model).Error; err != nil {
		fmt.Printf("DeleteById.Err:%s\n", err)
		return err
	}
	return nil
}

func Create(model interface{}) error {
	if err := dbClient.Create(model).Error; err != nil {
		fmt.Printf("Create.Err:%s\n", err)
		return err
	}
	return nil
}

func Update(model interface{}) error {
	if err := dbClient.Update(model).Error; err != nil {
		fmt.Printf("Update.Err:%s\n", err)
		return err
	}
	return nil
}

func ParseSearchMap(searchMap map[string]interface{}) (string, []interface{}) {
	var querySql string
	var index int = 0
	queryArgs := make([]interface{}, len(searchMap))
	for searchKey, searchValue := range searchMap {
		if len(searchKey) > 0 {
			switch vv := searchValue.(type) {
			case string:
				if len(vv) > 0 {
					if len(querySql) == 0 {
						querySql += " " + searchKey + " LIKE  ? "
					} else {
						querySql += " AND " + searchKey + " LIKE  ? "
					}
					queryArgs[index] = searchValue
				}
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
				if len(vv) > 0 {
					var stringSlice []string
					for _, u := range vv {
						stringSlice = append(stringSlice, u.(string))
					}
					values := strings.Join(stringSlice, ",")
					if len(values) > 0 {
						if len(querySql) == 0 {
							querySql += " " + searchKey + " in (?) "
						} else {
							querySql += " AND " + searchKey + " in (?) "
						}
						queryArgs[index] = values
					}
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
	if dbClient != nil {
		dbClient.Close()
	}
}
