package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/controller/vo/req"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"net/http"
	"strconv"
	"strings"
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
		idStr, _ := ctx.GetQuery("id")
		if len(idStr) <= 0 {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		roleId, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, service.GetRoleById(uint(roleId)))
	}
}

func CreateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}

func UpdateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}

func DeleteRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var deleteReq req.BaseDeleteReq
		error := ctx.BindJSON(&deleteReq)
		if error != nil {
			fmt.Printf("DeleteUser.bind.error.%s \n", error)
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}

		if len(deleteReq.Ids) > 0 {
			var roleIds []string = strings.Split(deleteReq.Ids, ",")
			for _, roleIdStr := range roleIds {
				roleId, err := strconv.Atoi(roleIdStr)
				if err == nil {
					service.DeleteRoleById(uint(roleId))
				}
			}
		}
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}
