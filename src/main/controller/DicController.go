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
		idStr, _ := ctx.GetQuery("id")
		if len(idStr) <= 0 {
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		dicId, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		ctx.JSON(http.StatusOK, service.GetDicById(uint(dicId)))
	}
}

func CreateDic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysDic model.SysDic
		error := ctx.BindJSON(&sysDic)
		if error != nil {
			fmt.Printf("CreateDic.bind.error.%s \n", error)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
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
			fmt.Printf("UpdateDic.bind.error.%s \n", error)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.UpdateDic(&sysDic)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func DeleteDic() gin.HandlerFunc {
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
			var dicIds []string = strings.Split(deleteReq.Ids, ",")
			for _, dicIdStr := range dicIds {
				dicId, err := strconv.Atoi(dicIdStr)
				if err == nil {
					service.DeleteDicById(uint(dicId))
				}
			}
		}
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}
