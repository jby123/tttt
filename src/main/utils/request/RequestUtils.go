package request

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/comm/exception"
	"goAdmin/src/main/controller/vo/req"
	"goAdmin/src/main/utils"
	"strconv"
	"strings"
)

func ParseRequestId(ctx *gin.Context) int {
	idStr, _ := ctx.GetQuery("id")
	if len(idStr) <= 0 {
		//抛出异常
		exception.BusinessException(utils.BUSINESS_ERROR, "数据参数有误", nil)
	}
	targetId, err := strconv.Atoi(idStr)
	if err != nil {
		//抛出异常
		exception.BusinessException(utils.BUSINESS_ERROR, "数据参数有误", nil)
	}
	return targetId
}

func ParseRequestIds(ctx *gin.Context) []int {
	var deleteReq req.BaseDeleteReq
	error := ctx.BindJSON(&deleteReq)
	if error != nil {
		exception.BusinessException(utils.BUSINESS_ERROR, "数据参数有误", nil)
	}
	return ParseStrToIds(deleteReq.Ids)
}

func ParseStrToIds(strResourceId string) []int {
	var idList []int
	if len(strResourceId) > 0 {
		var strIds []string = strings.Split(strResourceId, ",")
		idList = make([]int, len(strIds))
		var index int = 0
		for _, IdStr := range strIds {
			id, err := strconv.Atoi(IdStr)
			if err == nil {
				idList[index] = id
			}
			index++
		}
	}

	return idList

}
