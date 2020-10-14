package controllers

import (
	"GinDemo/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserListCtrl(c *gin.Context) {
	pageNum, pageSize := c.DefaultQuery("PageNum", "10"), c.DefaultQuery("PageSize", "1")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	var userDto service.UserDto
	res, err := userDto.ListUser(num, size, nil)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
	}
	c.JSON(200, gin.H{
		"message": res,
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
}
