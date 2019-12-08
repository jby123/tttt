package req

type MoveUserReq struct {
	DepartmentId uint   `json:"departmentId" binding:"required"`
	UserIds      []uint `json:"userIds" binding:"required"`
}
