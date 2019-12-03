package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func ResultSuccess(ctx *gin.Context, objects interface{}) {
	ctx.JSON(http.StatusOK, Success(objects))
}

func ResultError(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(http.StatusOK, Error(status, message, data))
}

func Success(objects interface{}) map[string]interface{} {
	return gin.H{
		"status": SuccessCode,
		"msg":    SuccessMessage,
		"data":   objects,
	}
}

func Error(status int, message string, data interface{}) map[string]interface{} {
	return gin.H{
		"status": status,
		"msg":    message,
		"data":   data,
	}
}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data

}
func transformer(objects interface{}) interface{} {
	getType := reflect.TypeOf(objects)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(objects)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.Elem().NumField(); i++ {
		field := getType.Elem().Field(i)
		value := getValue.Elem().Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

	return objects
}
