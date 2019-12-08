package controller

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/controller/vo/req"
	"goAdmin/src/main/model"
	"goAdmin/src/main/service"
	"goAdmin/src/main/utils"
	"net/http"
	"strconv"
)

/**
 * 分页获取日志列表
 */
func PageLogs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basePageReq := req.ParsePageReq(ctx)
		page := service.FindLogByPage(basePageReq.KeyWord, basePageReq.Order, basePageReq.Sort, basePageReq.Page, basePageReq.Size)
		ctx.JSON(http.StatusOK, utils.Success(page))
	}
}

func ClearLogs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service.DeleteLogsAll()
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func GetKeepLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keepLogDic := service.GetDicByName(utils.DicByLogKeepKey)
		if keepLogDic == nil || keepLogDic.ID <= 0 {
			keepLogDic = &model.SysDic{Name: utils.DicByLogKeepKey, Value: strconv.Itoa(utils.DicByLogKeepValue), Status: 1, Remark: "默认日志保存天数"}
			_ = service.CreateDic(keepLogDic)
		}
		//设置日志保存天数值
		ctx.JSON(http.StatusOK, utils.Success(keepLogDic.Value))
	}
}

func SetKeepLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var logSetKeepReq req.LogSetKeepReq
		error := ctx.BindJSON(&logSetKeepReq)
		if error != nil {
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		keepLogDic := service.GetDicByName(utils.DicByLogKeepKey)
		if keepLogDic == nil || keepLogDic.ID <= 0 {
			keepLogDic = &model.SysDic{Name: utils.DicByLogKeepKey, Value: strconv.Itoa(logSetKeepReq.Value), Status: 1, Remark: "默认日志保存天数"}
			_ = service.CreateDic(keepLogDic)
		} else {
			keepLogDic.Value = strconv.Itoa(logSetKeepReq.Value)
			_ = service.UpdateDic(keepLogDic)
		}
		//设置日志保存天数值
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}
