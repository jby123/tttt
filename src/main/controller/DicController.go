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
 * 分页获取-字典-列表
 */
func PageDic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basePageReq := req.ParsePageReq(ctx)
		page := service.FindDicByPage(basePageReq.KeyWord, basePageReq.Order, basePageReq.Sort, basePageReq.Page, basePageReq.Size)
		ctx.JSON(http.StatusOK, utils.Success(page))
	}
}

func GetDic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dicId := request.ParseRequestId(ctx)
		utils.ResultSuccess(ctx, service.GetDicById(uint(dicId)))
	}
}

func CreateDic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysDic model.SysDic
		error := ctx.BindJSON(&sysDic)
		if error != nil {
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.CreateDic(&sysDic)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func UpdateDic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysDic model.SysDic
		error := ctx.BindJSON(&sysDic)
		if error != nil {
			ctx.JSON(http.StatusOK, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.UpdateDic(&sysDic)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func DeleteDic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dicIds := request.ParseRequestIds(ctx)
		for _, dicId := range dicIds {
			service.DeleteDicById(uint(dicId))
		}
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}
