package controller

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/controller/vo/req"
	"goAdmin/src/main/model"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"goAdmin/src/main/utils/request"
	"net/http"
)

/**
 * 分页获取部门列表
 */
func PageDepartments() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basePageReq := req.ParsePageReq(ctx)
		page := service.FindDeptByPage(basePageReq.KeyWord, basePageReq.Order, basePageReq.Sort, basePageReq.Page, basePageReq.Size)
		ctx.JSON(http.StatusOK, utils.Success(page))
	}
}

func ListDepartments() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchMap := make(map[string]interface{})
		_, resultDataList := service.FindDeptListByParam(searchMap)
		ctx.JSON(http.StatusOK, utils.Success(resultDataList))
	}
}

func GetDepartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		deptId := request.ParseRequestId(ctx)
		utils.ResultSuccess(ctx, service.GetDeptById(uint(deptId)))
	}
}

func CreateDepartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysDept model.SysDepartment
		error := ctx.BindJSON(&sysDept)
		if error != nil {
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.CreateDept(&sysDept)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func UpdateDepartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysDept model.SysDepartment
		error := ctx.BindJSON(&sysDept)
		if error != nil {
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.UpdateDept(&sysDept)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func DeleteDepartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		deptIds := request.ParseRequestIds(ctx)
		for _, deptId := range deptIds {
			service.DeleteDeptById(uint(deptId))
		}
		ctx.JSON(http.StatusOK, nil)
	}
}
