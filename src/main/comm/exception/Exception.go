package exception

import "goAdmin/src/main/utils"

type Model struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	args    []interface{} `json:"args"`
}

/**
 * 业务异常
 */
func BusinessException(code int, message string, args ...interface{}) {
	if code == 0 {
		code = utils.BUSINESS_ERROR
	}
	model := Model{Code: code, Message: message, args: args}
	panic(model)
}

/**
 * 系统异常
 */
func SystemException(message string) {
	if len(message) == 0 {
		message = utils.SYSTEM_MESSAGE
	}
	model := Model{Code: utils.SYSTEM_ERROR_CODE, Message: message, args: nil}
	panic(model)
}

/**
 * 未授权异常
 */
func UnAuthorizationException() {
	model := Model{Code: 404, Message: "account.not.authorization"}
	panic(model)
}

/**
 * token过期异常
 */
func TokenExpiredException() {
	model := Model{Code: 404, Message: "account.not.authorization"}
	panic(model)
}

/**
 * 登入账号异常
 */
func LoginAccountException() {
	model := Model{Code: utils.BUSINESS_ERROR, Message: "account.or.password.exception"}
	panic(model)
}
