package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func GetError(dto interface{}, err validator.ValidationErrors) string {
	errInfo := err[0]
	refUser := reflect.ValueOf(dto)
	methodName := "Validate" + strings.Split(errInfo.Namespace(), ".")[1]
	method := refUser.MethodByName(methodName)
	args := []reflect.Value{reflect.ValueOf(errInfo)}
	res := method.Call(args)
	return res[0].String()
}
