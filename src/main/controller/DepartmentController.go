package controller

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/controller/vo/req"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"net/http"
	"strconv"
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
		idStr, _ := ctx.GetQuery("id")
		if len(idStr) <= 0 {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		deptId, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, service.GetDeptById(uint(deptId)))
	}
}

func CreateDepartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}

func UpdateDepartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}

func DeleteDepartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}
