package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xinliangnote/go-util/time"
	"goAdmin/src/main/utils"
	"log"
	"os"
)

/**
 *全局日志拦截中间件
 */
func LogHandler() gin.HandlerFunc {
	go handleAccessChannel()

	return func(context *gin.Context) {
		fmt.Println("before request....<<<<<LogHandler.begin>>>>")
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
		context.Writer = bodyLogWriter
		// 开始时间
		startTime := time.GetCurrentMilliUnix()
		// 处理请求
		context.Next()

		responseBody := bodyLogWriter.body.String()

		var responseCode int
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			res := utils.ResultData{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseCode = res.Status
				responseMsg = res.Msg
				responseData = res.Data
			}
		}

		// 结束时间
		endTime := time.GetCurrentMilliUnix()

		if context.Request.Method == "POST" {
			_ = context.Request.ParseForm()
		}

		// 日志格式
		accessLogMap := make(map[string]interface{})

		accessLogMap["request_time"] = startTime
		accessLogMap["request_method"] = context.Request.Method
		accessLogMap["request_uri"] = context.Request.RequestURI
		accessLogMap["request_proto"] = context.Request.Proto
		accessLogMap["request_ua"] = context.Request.UserAgent()
		accessLogMap["request_referer"] = context.Request.Referer()
		accessLogMap["request_post_data"] = context.Request.PostForm.Encode()
		accessLogMap["request_client_ip"] = context.ClientIP()

		accessLogMap["response_time"] = endTime
		accessLogMap["response_code"] = responseCode
		accessLogMap["response_msg"] = responseMsg
		accessLogMap["response_data"] = responseData

		accessLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)

		accessLogJson, _ := utils.Encode(accessLogMap)
		accessChannel <- accessLogJson

		fmt.Printf("\n<<<<<<after request ....<<<<LogHandler.end.\n...json: %s>>>>>", accessLogJson)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var accessChannel = make(chan string, 100)

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func handleAccessChannel() {
	if f, err := os.OpenFile(utils.AppAccessLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		log.Println(err)
	} else {
		for accessLog := range accessChannel {
			_, _ = f.WriteString(accessLog + "\n")
		}
	}
	return
}
