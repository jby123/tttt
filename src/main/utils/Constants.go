package utils

import "time"

const AppName = "goAdmin"
const DefaultDevelopmentRelativePath = "./src/resources"
const DefaultTestRelativePath = "E:\\go_project\\goAdmin\\src\\test\\resources"
const DefaultDevelopmentEnv = "dev"

const DefaultPassword = "123456"
const DefaultCaptchaLen = 4

const DEFAULT_EXPIRE_SECOND_TIME = time.Duration(30) * time.Second

const DEFAULT_EXPIRE_HOURE_TIME = time.Duration(1) * time.Hour

const DEFAULT_EXPIRE_DAY_TIME = time.Duration(24) * time.Hour

const SuccessCode int = 1000
const SYSTEM_ERROR_CODE int = 500
const SYSTEM_BUSSINESS_CODE = 400
const BUSINESS_ERROR = 1001
const SuccessMessage string = "SUCCESS"
const SYSTEM_MESSAGE string = "UNKNOWN.ERROR"

// 日志文件
const AppAccessLogName = "log/" + AppName + "-access.log"

const DicByLogKeepKey = "LOG_KEEP_SAVE_DAY_KEY"

const DicByLogKeepValue int = 30 //日志默认保存30天

//缓存数据参数
const (
	CacheIpKey   = "biz:ip:all"
	CacheUserKey = "biz_user:%d"
)
