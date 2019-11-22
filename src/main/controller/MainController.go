package controller

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/utils"
	"net/http"
)

func Logout(ctx gin.Context) {
	ctx.Status(http.StatusOK)
	ctx.JSON(utils.SuccessCode, nil)
}
