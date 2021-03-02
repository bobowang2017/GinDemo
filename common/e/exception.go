package e

import "net/http"

const (
	SERVER_ERROR    = 1000 // 系统错误
	NOT_FOUND       = 1001 // 401错误
	UNKNOWN_ERROR   = 1002 // 未知错误
	PARAMETER_ERROR = 1003 // 参数错误
)

// api错误的结构体
type APIException struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.Msg
}

func newAPIException(status int, msg string) *APIException {
	return &APIException{
		Status: status,
		Msg:    msg,
	}
}

// 500 错误处理
func ServerError(message string) *APIException {
	return newAPIException(SERVER_ERROR, message)
}

// 404 错误
func NotFound(message string) *APIException {
	return newAPIException(NOT_FOUND, http.StatusText(http.StatusNotFound))
}

// 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(UNKNOWN_ERROR, message)
}

// 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(PARAMETER_ERROR, message)
}
