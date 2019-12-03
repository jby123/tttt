package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/controller/vo/req"
	"goAdmin/src/main/model"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"goAdmin/src/main/utils/request"
	"net/http"
)

/**
 * 分页获取角色列表
 */
func PageRoles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basePageReq := req.ParsePageReq(ctx)
		page := service.FindRoleByPage(basePageReq.KeyWord, basePageReq.Order, basePageReq.Sort, basePageReq.Page, basePageReq.Size)
		ctx.JSON(http.StatusOK, utils.Success(page))
	}
}

func ListRoles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchMap := make(map[string]interface{})
		_, resultDataList := service.FindRoleListByParam(searchMap)
		ctx.JSON(http.StatusOK, utils.Success(resultDataList))
	}
}

func GetRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleId := request.ParseRequestId(ctx)
		ctx.JSON(http.StatusOK, service.GetRoleById(uint(roleId)))
	}
}

func CreateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysRole model.SysRole
		error := ctx.BindJSON(&sysRole)
		if error != nil {
			fmt.Printf("CreateRole.bind.error.%s \n", error)
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.CreateRole(&sysRole)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func UpdateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysRole model.SysRole
		error := ctx.BindJSON(&sysRole)
		if error != nil {
			fmt.Printf("CreateRole.bind.error.%s \n", error)
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.UpdateRole(&sysRole)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func DeleteRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleIds := request.ParseRequestIds(ctx)
		for _, roleId := range roleIds {
			service.DeleteRoleById(uint(roleId))
		}
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}
