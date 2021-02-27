package controllers

import (
	"GinDemo/common/e"
	"GinDemo/dto"
	"GinDemo/service"
	"GinDemo/utils/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func UserRouterRegister(router *gin.RouterGroup) {
	user := UserController{
		service.NewUserService(),
	}
	router.GET("/", e.Wrapper(user.UserList))
	router.POST("/", e.Wrapper(user.UserAdd))
	router.DELETE("/:userId", e.Wrapper(user.UserDel))
	router.PUT("/:userId", e.Wrapper(user.UserUpdate))
	router.GET("/test", e.Wrapper(user.UserTest))
}

func (u *UserController) UserList(c *gin.Context) interface{} {
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

	res, err := u.userService.ListUser(num, size, params)
	if err != nil {
		logger.Error(err)
		return e.ServerError()
	}
	return res
}

func (u *UserController) UserAdd(c *gin.Context) interface{} {
	if err := c.BindJSON(&dto.UserDto{}); err != nil {
		return e.ParameterError(err.Error())
	}
	return "success"
}

func (u *UserController) UserDel(c *gin.Context) interface{} {
	userId, _ := strconv.Atoi(c.Param("userId"))
	err := u.userService.DelByUserId(userId)
	if err != nil {
		logger.Error(err)
		return e.ServerError()
	}
	return "success"
}

func (u *UserController) UserUpdate(c *gin.Context) interface{} {
	userId, _ := strconv.Atoi(c.Param("userId"))
	params := map[string]interface{}{
		"name": "wangxiangbo",
		"sex":  88,
	}
	err := u.userService.UpdateByUserId(userId, params)
	if err != nil {
		logger.Error(err)
		return e.ServerError()
	}
	return "success"
}

func (u *UserController) UserTest(c *gin.Context) interface{} {
	Username := c.Query("Username")
	param := map[string]interface{}{
		"Username": Username,
	}
	res, err := u.userService.SearchUser(param)
	if err != nil {
		return e.ServerError()
	}
	return res
}
