package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/controller/vo/req"
	"goAdmin/src/main/model"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"net/http"
	"strconv"
	"strings"
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
		idStr, _ := ctx.GetQuery("id")
		if len(idStr) <= 0 {
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		paramId, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		ctx.JSON(http.StatusOK, service.GetParamById(uint(paramId)))
	}
}

func CreateParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysParam model.SysParam
		error := ctx.BindJSON(&sysParam)
		if error != nil {
			fmt.Printf("CreateParam.bind.error.%s \n", error)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
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
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.UpdateParam(&sysParam)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func DeleteParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var deleteReq req.BaseDeleteReq
		error := ctx.BindJSON(&deleteReq)
		if error != nil {
			fmt.Printf("DeleteParam.bind.error.%s \n", error)
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}

		if len(deleteReq.Ids) > 0 {
			var paramIds []string = strings.Split(deleteReq.Ids, ",")
			for _, paramIdStr := range paramIds {
				paramId, err := strconv.Atoi(paramIdStr)
				if err == nil {
					service.DeleteParamById(uint(paramId))
				}
			}
		}
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}
