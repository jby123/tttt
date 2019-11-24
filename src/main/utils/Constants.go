package utils

import "time"

const AppName = "goAdmin"
const SigningKey = "tehabi@admin.com"
const Issuer = "lucky" //签名发行者
const DefaultDevelopmentRelativePath = "./src/resources"
const DefaultTestRelativePath = "./src/test/resources"
const DefaultDevelopmentEnv = "dev"

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
