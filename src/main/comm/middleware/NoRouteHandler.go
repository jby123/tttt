package middleware

import (
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/utils"
	"net/http"
)

func GlobalNoRouteHandler(context *gin.Context) {
	context.JSON(http.StatusNotFound, utils.Error(http.StatusBadRequest, "NOT.FOUND", nil))
	return
}
