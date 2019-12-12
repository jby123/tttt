package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rocket049/gostructcopy"
	"goAdmin/src/main/comm/exception"
	"goAdmin/src/main/controller/vo"
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
		resultDataList := service.FindRoleListByParam(searchMap)
		ctx.JSON(http.StatusOK, utils.Success(resultDataList))
	}
}

func GetRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleId := request.ParseRequestId(ctx)
		roleInfo := service.GetRoleById(uint(roleId))

		if roleInfo != nil {
			menuIds, _ := service.FindMenuIdListByRoleId(uint(roleId))
			deptIds, _ := service.FindDeptIdListByRoleId(uint(roleId))
			targetRoleInfo := vo.SysRoleVo{
				ID:               roleInfo.ID,
				Status:           roleInfo.Status,
				Name:             roleInfo.Name,
				Remark:           roleInfo.Remark,
				Code:             roleInfo.Code,
				MenuIdList:       menuIds,
				DepartmentIdList: deptIds,
			}
			utils.ResultSuccess(ctx, targetRoleInfo)
			return
		}
		utils.ResultSuccess(ctx, nil)
	}
}

func CreateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var srcData req.SysRoleReq
		error := ctx.BindJSON(&srcData)
		if error != nil {
			fmt.Printf("CreateRole.bind.error.%s \n", error)
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		targetData := new(model.SysRole)
		err := gostructcopy.StructCopy(srcData, targetData)
		if err != nil {
			exception.SystemException(err.Error())
		}

		//保存角色信息
		err = service.CreateRole(targetData)
		if err != nil || targetData.ID <= 0 {
			exception.SystemException("保存角色信息失败")
		}
		//保存角色和部门的关系
		service.SaveOrUpdateRoleDeptBatch(targetData.ID, srcData.DepartmentIdList)
		//保存角色对应的菜单
		service.SaveOrUpdateRoleMenuBatch(targetData.ID, srcData.MenuIdList)

		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func UpdateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var srcData req.SysRoleReq
		error := ctx.BindJSON(&srcData)
		if error != nil {
			fmt.Printf("CreateRole.bind.error.%s \n", error)
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}

		targetData := new(model.SysRole)
		err := gostructcopy.StructCopy(srcData, targetData)
		if err != nil {
			exception.SystemException(err.Error())
		}
		//修改角色信息
		err = service.UpdateRole(targetData)
		if err != nil || targetData.ID <= 0 {
			exception.SystemException("修改角色信息失败" + err.Error())
		}
		//保存角色和部门的关系
		service.SaveOrUpdateRoleDeptBatch(targetData.ID, srcData.DepartmentIdList)
		//保存角色对应的菜单
		service.SaveOrUpdateRoleMenuBatch(targetData.ID, srcData.MenuIdList)

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
