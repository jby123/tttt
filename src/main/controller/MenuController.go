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
 * 分页获取菜单列表
 */
func PageMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basePageReq := req.ParsePageReq(ctx)
		page := service.FindMenuByPage(basePageReq.KeyWord, basePageReq.Order, basePageReq.Sort, basePageReq.Page, basePageReq.Size)
		ctx.JSON(http.StatusOK, utils.Success(page))
	}
}

func ListMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchMap := make(map[string]interface{})
		_, resultDataList := service.FindMenuListByParam(searchMap)
		ctx.JSON(http.StatusOK, utils.Success(resultDataList))
	}
}
func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr, _ := ctx.GetQuery("id")
		if len(idStr) <= 0 {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		menuId, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, service.GetMenuById(uint(menuId)))
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysMenu model.SysMenu
		error := ctx.BindJSON(&sysMenu)
		if error != nil {
			fmt.Printf("UpdateMenu.bind.error.%s \n", error)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.CreateMenu(&sysMenu)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sysMenu model.SysMenu
		error := ctx.BindJSON(&sysMenu)
		if error != nil {
			fmt.Printf("UpdateMenu.bind.error.%s \n", error)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}
		service.UpdateMenu(&sysMenu)
		ctx.JSON(http.StatusOK, utils.Success(nil))
	}
}

func DeleteMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var deleteReq req.BaseDeleteReq
		error := ctx.BindJSON(&deleteReq)
		if error != nil {
			fmt.Printf("DeleteMenu.bind.error.%s \n", error)
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, utils.Error(utils.BUSINESS_ERROR, "数据参数有误", nil))
			return
		}

		if len(deleteReq.Ids) > 0 {
			var menuIds []string = strings.Split(deleteReq.Ids, ",")
			for _, menuIdStr := range menuIds {
				menuId, err := strconv.Atoi(menuIdStr)
				if err == nil {
					service.DeleteMenuById(uint(menuId))
				}
			}
		}
		ctx.Status(http.StatusOK)
		ctx.JSON(http.StatusOK, nil)
	}
}
