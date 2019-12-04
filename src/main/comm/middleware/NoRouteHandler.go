package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/src/main/utils"
	"net/http"
)

func GlobalNoRouteHandler(context *gin.Context) {
	fmt.Println("GlobalNoRouteHandler request.." + context.Request.RequestURI + "..<<<<<GlobalNoRouteHandler.end>>>>")
	context.JSON(http.StatusNotFound, utils.Error(http.StatusBadRequest, "NOT.FOUND", nil))
	return
}
