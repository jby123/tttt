package req

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type BasePageReq struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	Sort    string `json:"sort"`  //desc
	Order   string `json:"order"` //排序的字段
	Time    int    `json:"t"`
	KeyWord string `json:"keyWord"`
}

func ParsePageReq(ctx *gin.Context) BasePageReq {
	keyWord := ctx.Query("keyWord")
	pageStr := ctx.Query("page")
	sizeStr := ctx.Query("size")
	timeStr := ctx.Query("t")
	sort := ctx.Query("sort")
	order := ctx.Query("order")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	timeReq, err := strconv.Atoi(timeStr)
	req := BasePageReq{
		Page:    page,
		Size:    size,
		Sort:    sort,
		Order:   order,
		Time:    timeReq,
		KeyWord: keyWord,
	}
	return req
}
