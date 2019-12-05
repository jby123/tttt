package req

type SysRoleReq struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Code             string `json:"label"`
	Status           uint   `json:"status"`
	Remark           string `json:"remark"`
	MenuIdList       []uint `json:"menuIdList"`
	DepartmentIdList []uint `json:"departmentIdList"`
}
