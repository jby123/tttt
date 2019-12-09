package database

type PaginationVo struct {
	List interface{} `json:"list"`
	Page *PageInfo   `json:"pagination"`
}

type PageInfo struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

/**
 *分页请求对象封装 [针对单表]
 */
type Page struct {
	SearchMap      map[string]interface{} //查询条件 interface 可以是直接的值 也可以是 Condition条件
	Model          interface{}            //查询表的实体model
	resultDataList interface{}            //返回结果集
	Page           *PageInfo              //分页请求参数
}

func Pagination(resultDataList interface{}, pageNum, pageSize, total int) *PaginationVo {
	paginationVo := &PaginationVo{List: resultDataList, Page: &PageInfo{Page: pageNum, Size: pageSize, Total: total}}
	return paginationVo
}
