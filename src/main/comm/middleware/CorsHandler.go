package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsHandler() gin.HandlerFunc {
	return func(app *gin.Context) {
		app.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		app.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		app.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//放行所有OPTIONS方法
		method := app.Request.Method
		if method == "OPTIONS" {
			app.JSON(http.StatusOK, "Options Request!")
			return
		}
		app.Next()
	}
}
