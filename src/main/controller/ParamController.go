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
 * 分页获取-系統參數-列表
 */
func PageParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basePageReq := req.ParsePageReq(ctx)
		page := service.FindParamByPage(basePageReq.KeyWord, basePageReq.Order, basePageReq.Sort, basePageReq.Page, basePageReq.Size)
		ctx.JSON(http.StatusOK, utils.Success(page))
	}
}

func GetParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramId := request.ParseRequestId(ctx)
		utils.ResultSuccess(ctx, service.GetParamById(uint(paramId)))
	}
}

func CreateParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysParam model.SysParam
		error := ctx.BindJSON(&sysParam)
		if error != nil {
			fmt.Printf("CreateParam.bind.error.%s \n", error)
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.CreateParam(&sysParam)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func UpdateParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysParam model.SysParam
		error := ctx.BindJSON(&sysParam)
		if error != nil {
			fmt.Printf("UpdateParam.bind.error.%s \n", error)
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.UpdateParam(&sysParam)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func DeleteParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramIds := request.ParseRequestIds(ctx)
		for _, paramId := range paramIds {
			service.DeleteParamById(uint(paramId))
		}
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}
