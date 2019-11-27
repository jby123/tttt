package controller

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/controller/vo/req"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
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