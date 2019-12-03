package database

type PaginationVo struct {
	List interface{} `json:"list"`
	Page PageInfo    `json:"pagination"`
}

type PageInfo struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

func Pagination(resultDataList interface{}, pageNum, pageSize, total int) *PaginationVo {
	paginationVo := &PaginationVo{List: resultDataList, Page: PageInfo{Page: pageNum, Size: pageSize, Total: total}}
	return paginationVo
}
