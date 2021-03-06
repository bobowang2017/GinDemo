package dto

import (
	"github.com/go-playground/validator/v10"
)

type UserDto struct {
	PageNum  int
	PageSize int
	Username string `binding:"required"`
	Password string `binding:"required,min=6,max=12"`
	Name     string
	Sex      *int `binding:"required"`
	Birthday string
}

func (r *UserDto) ValidatePassword(errInfo validator.FieldError) string {
	switch errInfo.Tag() {
	case "min":
		return "密码长度不能小于6位"
	case "max":
		return "密码长度不能超过12位"
	case "required":
		return "密码不能为空"
	default:
		return "密码验证失败"
	}
}

func (r *UserDto) ValidateUsername(errInfo validator.FieldError) string {
	switch errInfo.Tag() {
	case "required":
		return "用户名不能为空"
	default:
		return "用户名验证失败"
	}
}
