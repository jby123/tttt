package req

type SysUserReq struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	NickName     string `json:"nickName"`
	Password     string `json:"password"` //未加密参数
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	HeadImg      string `json:"headImg"`
	DepartmentId uint   `json:"departmentId"`
	Status       uint   `json:"status"`
	Remark       string `json:"remark"`
	RoleId       uint   `json:"roleId"`
}
