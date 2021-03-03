package controllers

import (
	"GinDemo/common/e"
	"GinDemo/dto"
	"GinDemo/models"
	"GinDemo/service"
	"GinDemo/utils"
	"GinDemo/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

type UserController struct {
	userService service.IUserService
}

func UserRouterRegister(router *gin.RouterGroup) {
	user := UserController{
		service.NewUserService(),
	}
	router.GET("/", e.Wrapper(user.UserList))
	router.POST("/", e.Wrapper(user.UserAdd))
	router.DELETE("/:userId", e.Wrapper(user.UserDel))
	router.PUT("/:userId", e.Wrapper(user.UserUpdate))
	router.GET("/:userId", e.Wrapper(user.UserDetail))
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
		return e.ServerError(err.Error())
	}
	return res
}

func (u *UserController) UserAdd(c *gin.Context) interface{} {
	uDto := &dto.UserDto{}
	if err := c.BindJSON(uDto); err != nil {
		return e.ParameterError(uDto.GetError(uDto, err.(validator.ValidationErrors)))
	}
	user := &models.User{
		Username: uDto.Username,
		Password: uDto.Password,
		Name:     uDto.Name,
		Sex:      uDto.Sex,
	}
	if uDto.Birthday != "" {
		bir, err := time.ParseInLocation("2006-01-02", uDto.Birthday, time.Local)
		if err != nil {
			return e.ParameterError(err.Error())
		}
		user.Birthday = &utils.JSONTime{Time: bir}
	}

	err := u.userService.CreateUser(user)
	if err != nil {
		return e.ParameterError(err.Error())
	}
	return "success"
}

func (u *UserController) UserDel(c *gin.Context) interface{} {
	userId, _ := strconv.Atoi(c.Param("userId"))
	err := u.userService.DelByUserId(userId)
	if err != nil {
		logger.Error(err)
		return e.ServerError(err.Error())
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
		return e.ServerError(err.Error())
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
		return e.ServerError(err.Error())
	}
	return res
}

func (u *UserController) UserDetail(c *gin.Context) interface{} {
	userId, _ := strconv.Atoi(c.Param("userId"))
	user, err := u.userService.Detail(userId)
	if err != nil {
		logger.Error(err)
		return e.ServerError(err.Error())
	}
	return user
}
