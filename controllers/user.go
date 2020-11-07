package controllers

import (
	"GinDemo/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserListCtrl(c *gin.Context) {
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
		c.JSON(200, gin.H{
			"message": err,
		})
	}
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "ok",
		"data":   res,
	})
}

func UserAddCtrl(c *gin.Context) {
	var userDto service.UserDto
	c.BindJSON(&userDto)
	userDto.CreateUser()
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func UserDelCtrl(c *gin.Context) {
	var userDto service.UserDto
	userId, _ := strconv.Atoi(c.Param("userId"))
	userDto.DelByUserId(userId)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func UserUpdateCtrl(c *gin.Context) {
	var userDto service.UserDto
	userId, _ := strconv.Atoi(c.Param("userId"))
	params := map[string]interface{}{
		"name": "wangxiangbo",
		"sex":  88,
	}
	userDto.UpdateByUserId(userId, params)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func UserTestCtrl(c *gin.Context) {
	Username := c.Query("Username")
	var userDto service.UserDto
	param := map[string]interface{}{
		"Username": Username,
	}
	res, err := userDto.SearchUser(param)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "ok",
			"data":    res,
		})
	}

}
