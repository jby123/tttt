package utils

import "time"

const DEFAULT_EXPIRE_SECOND_TIME  = time.Duration(30)*time.Second

const DEFAULT_EXPIRE_HOURE_TIME  = time.Duration(1)*time.Hour

const DEFAULT_EXPIRE_DAY_TIME  = time.Duration(24)*time.Hour

const SUCCESS_CODE  int = 200
const SYSTEM_ERROR_CODE int  = 500
const SYSTEM_BUSSINESS_CODE  = 400
const SUCCESS_MESSAGE string = "SUCCESS"
const SYSTEM_MESSAGE string = "UNKNOWN.ERROR"
