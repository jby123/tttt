package controller

import (
	"github.com/gin-gonic/gin"
	utils2 "goAdmin/src/main/utils"
	"net/http"
)

func Login(ctx gin.Context) {
	ctx.Status(http.StatusOK)
	ctx.JSON(utils.SUCCESS_CODE,nil)
}

