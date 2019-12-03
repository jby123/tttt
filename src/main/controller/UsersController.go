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
* @api /user/page 分页获取用户列表
* @apiGroup user
* @apiVersion 1.0
* @apiDescription 获取所有的账号
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func PageUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basePageReq := req.ParsePageReq(ctx)
		var deptIdList []int
		if departmentIdsStr, isExist := ctx.GetQuery("departmentIds"); isExist {
			deptIdList = request.ParseStrToIds(departmentIdsStr)
		}
		page := service.FindUserByPage(deptIdList, basePageReq.KeyWord, basePageReq.Order, basePageReq.Sort, basePageReq.Page, basePageReq.Size)
		ctx.JSON(http.StatusOK, utils.Success(page))
	}
}

func ListUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

/**
* @api /user/info 根据id获取用户信息
* @apiGroup user
* @apiVersion 1.0
* @apiDescription 根据id获取用户信息
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := request.ParseRequestId(ctx)
		ctx.JSON(http.StatusOK, service.GetUserById(uint(userId)))
	}
}

/**
* @api /user/add 新建账号
* @apiGroup user
* @apiVersion 1.0
* @apiDescription 新建账号
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysUser model.SysUser
		error := ctx.BindJSON(&sysUser)
		if error != nil {
			fmt.Printf("CreateUser.bind.error.%s \n", error)
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.CreateUser(&sysUser)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

/**
* @api /user/update 更新账号
* @apiGroup user
* @apiVersion 1.0
* @apiDescription 更新账号
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysUser model.SysUser
		error := ctx.BindJSON(&sysUser)
		if error != nil {
			fmt.Printf("CreateUser.bind.error.%s \n", error)
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.UpdateUser(&sysUser)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

/**
* @api /user/delete 删除用户
* @apiGroup user
* @apiVersion 1.0
* @apiDescription 删除用户
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ids := request.ParseRequestIds(ctx)
		service.DeleteUserByIds(ids)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}
