package controllers

import (
	"GinDemo/common/e"
	"GinDemo/service"
	"GinDemo/utils/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserListCtrl(c *gin.Context) interface{} {
	pageNum := c.DefaultQuery("PageNum", "1")
	pageSize := c.DefaultQuery("PageSize", "10")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	Username := c.Query("Username")
	name := c.Query("name")
	params := make(map[string]interface{})
	if Username != "" {
		params["Username"] = Username
	}
	if name != "" {
		params["name"] = name
	}
	var userDto service.UserDto
	res, err := userDto.ListUser(num, size, params)
	if err != nil {
		logger.Error(err)
		return e.ServerError()
	}
	return res
}

func UserAddCtrl(c *gin.Context) interface{} {
	var userDto service.UserDto
	if err := c.BindJSON(&userDto); err != nil {
		return e.ParameterError(err.Error())
	}
	err := userDto.CreateUser()
	if err != nil {
		logger.Error(err)
		return e.ServerError()
	}
	return "success"
}

func UserDelCtrl(c *gin.Context) interface{} {
	var userDto service.UserDto
	userId, _ := strconv.Atoi(c.Param("userId"))
	err := userDto.DelByUserId(userId)
	if err != nil {
		logger.Error(err)
		return e.ServerError()
	}
	return "success"
}

func UserUpdateCtrl(c *gin.Context) interface{} {
	var userDto service.UserDto
	userId, _ := strconv.Atoi(c.Param("userId"))
	params := map[string]interface{}{
		"name": "wangxiangbo",
		"sex":  88,
	}
	err := userDto.UpdateByUserId(userId, params)
	if err != nil {
		logger.Error(err)
		return e.ServerError()
	}
	return "success"
}

func UserTestCtrl(c *gin.Context) interface{} {
	Username := c.Query("Username")
	var userDto service.UserDto
	param := map[string]interface{}{
		"Username": Username,
	}
	res, err := userDto.SearchUser(param)
	if err != nil {
		return e.ServerError()
	}
	return res
}
