package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goAdmin/src/main/comm/exception"
	"os"
	"time"
)

var dbClient *gorm.DB

const DefaultDevelopmentEnv = "dev"

//数据库常量
type DbConfig struct {
	Db MysqlConf `yaml:"DB"`
}

type MysqlConf struct {
	AutoCreateTable bool   `yaml:"AutoCreateTable"`
	DriverName      string `yaml:"DriverName"`
	Url             string `yaml:"Url"`
	Username        string `yaml:"Username"`
	Password        string `yaml:"Password"`
	Dialect         string `yaml:"Dialect"`
	MaxIdle         int    `yaml:"maxIdle"`
	MaxOpen         int    `yaml:"maxOpen"`
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
	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数

	//设置闲置的连接数
	dbClient.DB().SetMaxIdleConns(dbConfig.Db.MaxIdle)
	//设置最大打开的连接数
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

func FindCommPage(page *Page) *PaginationVo {

	err := FindPage(page.SearchMap, page.Pagination.Order, page.Pagination.Sort, page.Pagination.Page, page.Pagination.Size).Find(page.ResultDataList).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Pagination(make([]interface{}, 0), page.Pagination.Page, page.Pagination.Size, 0)
		}
		fmt.Println("findCommPage.error....>>>", err)
		return Pagination(make([]interface{}, 0), page.Pagination.Page, page.Pagination.Size, 0)
	}

	total := Count(page.SearchMap, page.Model)
	if total <= 0 {
		return Pagination(make([]interface{}, 0), page.Pagination.Page, page.Pagination.Size, 0)
	}
	return Pagination(page.ResultDataList, page.Pagination.Page, page.Pagination.Size, total)
}

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
	client := GetDB()
	searchSql, searchArgs := ParseSearchMap(searchMap)
	if len(searchSql) > 0 {
		client = client.Where(searchSql, searchArgs...)
	}
	if len(sort) <= 0 {
		sort = "DESC"
	}
	if len(order) <= 0 {
		order = "create_time"
	} else {
		order = CamelCase(order)
	}
	client = client.Order(" " + order + " " + sort + " ")
	if offset > 0 {
		client = client.Offset((offset - 1) * limit)
	}
	if limit > 0 {
		client = client.Limit(limit)
	}
	return client
}

func FindListByParam(searchMap map[string]interface{}, resultDataList interface{}) error {
	searchSql, searchArgs := ParseSearchMap(searchMap)
	client := GetDB()
	if len(searchSql) > 0 {
		client = client.Where(searchSql, searchArgs...)
	}
	if err := client.Find(resultDataList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// do something
			return nil
		}
		return err
	}
	return nil
}

func Count(searchMap map[string]interface{}, model interface{}) int {
	client := GetDB()
	client = client.Model(model)
	searchSql, searchArgs := ParseSearchMap(searchMap)
	if len(searchSql) > 0 {
		client = client.Where(searchSql, searchArgs...)
	}

	var total int = 0
	client.Count(&total)

	return total
}

/**
 * 通过 id 获取记录
 * @method GetById
 * @param  {[type]} model interface{} [description]
 */
func GetById(id uint, model interface{}) (error, bool) {
	client := GetDB()
	if err := client.Where("id = ?", id).First(model).Error; err != nil {
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, false
			}
			exception.SystemException(err.Error())
		}
	}
	return nil, true
}

func DeleteById(id uint, model interface{}) error {
	client := GetDB()
	if err := client.Where("id = ?", id).Delete(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// do something
			return nil
		}
		return err
	}
	return nil
}

func DeleteByIds(ids []int, model interface{}) error {
	client := GetDB()
	if err := client.Where("id in (?)", ids).Delete(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// do something
			return nil
		}
		return err
	}
	return nil
}
func DeleteByParams(searchMap map[string]interface{}, model interface{}) error {
	client := GetDB()
	client = client.Model(model)
	searchSql, searchArgs := ParseSearchMap(searchMap)
	if len(searchSql) > 0 {
		client = client.Where(searchSql, searchArgs...)
	}
	if err := client.Delete(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// do something
			return nil
		}
		return err
	}
	return nil
}

func Create(model interface{}) error {
	client := GetDB()
	if err := client.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func Update(model interface{}) error {
	client := GetDB()
	client = client.Model(model)
	if err := client.Update(model).Error; err != nil {
		return err
	}
	return nil
}

func UpdateById(id uint, model interface{}) error {
	client := GetDB()
	client = client.Model(model)
	if err := client.Where("id = ?", id).Update(model).Error; err != nil {
		return err
	}
	return nil
}

func ParseSearchMap(searchMap map[string]interface{}) (string, []interface{}) {
	var querySql string
	var index int = 0
	queryArgs := make([]interface{}, 0)
	for searchKey, searchValue := range searchMap {
		if len(searchKey) > 0 {
			switch vv := searchValue.(type) {
			case Condition:
				if len(querySql) == 0 {
					querySql += " " + searchKey
				} else {
					querySql += " AND " + searchKey
				}
				if vv.Operate == Equals {
					querySql += " =  ? "
					queryArgs = append(queryArgs, vv.Value)
				} else if vv.Operate == Like {
					querySql += " LIKE  ? "

					queryArgs = append(queryArgs, "%"+vv.Value+"%")
				} else if vv.Operate == LeftLike {
					querySql += " LIKE  ? "
					queryArgs = append(queryArgs, "%"+vv.Value+"%")
				} else if vv.Operate == RightLike {
					querySql += " LIKE  ? "
					queryArgs = append(queryArgs, ""+vv.Value+"%")
				} else {
					querySql += " LIKE  ? "
					queryArgs = append(queryArgs, "%"+vv.Value+"%")
				}

			case string:
				if len(vv) > 0 {
					if len(querySql) == 0 {
						querySql += " " + searchKey + " LIKE  ? "
					} else {
						querySql += " AND " + searchKey + " LIKE  ? "
					}
					queryArgs = append(queryArgs, "%"+searchValue.(string)+"%")
				}
			case float64:
				if len(querySql) == 0 {
					querySql += " " + searchKey + " =  ? "
				} else {
					querySql += " AND " + searchKey + " =  ? "
				}
				queryArgs = append(queryArgs, searchValue)
			case int:
				if len(querySql) == 0 {
					querySql += " " + searchKey + " =  ? "
				} else {
					querySql += " AND " + searchKey + " =  ? "
				}
				queryArgs = append(queryArgs, searchValue)
			case uint:
				if len(querySql) == 0 {
					querySql += " " + searchKey + " =  ? "
				} else {
					querySql += " AND " + searchKey + " =  ? "
				}
				queryArgs = append(queryArgs, searchValue)
			case []uint:
				if len(vv) > 0 {
					if len(querySql) == 0 {
						querySql += " " + searchKey + " in (?) "
					} else {
						querySql += " AND " + searchKey + " in (?) "
					}
					queryArgs = append(queryArgs, vv)
				}
			case []int:
				if len(vv) > 0 {
					if len(querySql) == 0 {
						querySql += " " + searchKey + " in (?) "
					} else {
						querySql += " AND " + searchKey + " in (?) "
					}
					queryArgs = append(queryArgs, vv)
				}
			case []interface{}:
				if len(vv) > 0 {
					if len(querySql) == 0 {
						querySql += " " + searchKey + " in (?) "
					} else {
						querySql += " AND " + searchKey + " in (?) "
					}
					queryArgs = append(queryArgs, vv)
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
