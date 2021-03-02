package e

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type HandlerFunc func(c *gin.Context) interface{}

func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		res := handler(c)
		resType := reflect.TypeOf(res).String()
		if resType == "*e.APIException" {
			var apiException *APIException
			if h, ok := res.(*APIException); ok {
				apiException = h
			} else if e, ok := res.(error); ok {
				if gin.Mode() == "debug" {
					// 错误
					apiException = UnknownError(e.Error())
				} else {
					// 未知错误
					apiException = UnknownError(e.Error())
				}
			} else {
				apiException = ServerError(e.Error())
			}
			c.JSON(200, apiException)
		} else {
			c.JSON(200, gin.H{
				"status": 200,
				"msg":    "ok",
				"data":   res,
			})
		}
	}
}
