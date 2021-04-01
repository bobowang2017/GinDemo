package controllers

import (
	"GinDemo/common/e"
	"GinDemo/common/utils"
	"GinDemo/common/utils/logger"
	"GinDemo/internal/dto"
	"GinDemo/internal/models"
	"GinDemo/internal/service"
	"github.com/gin-gonic/gin"
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

// @Title 用户
// @Author wangxiangbo
// @Description 查询用户列表
// @Tags User
// @Param operator query string false "操作者"
// @Param description query string false "操作描述"
// @Param start_time query string false "开始时间"
// @Param end_time query string false "结束时间"
// @Success 200
// @Router	/api/v1/users [get]
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
		return e.ParameterError(err.Error())
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

type UserAddrCtrl struct {
	userAddrService service.IUserAddrService
}

func UserAddressRouterRegister(router *gin.RouterGroup) {
	userAddr := UserAddrCtrl{
		service.NewUserAddressService(),
	}
	router.GET("/:userId/addresses", e.Wrapper(userAddr.UserAddrList))
	router.POST("/:userId/addresses", e.Wrapper(userAddr.UserAddrAdd))

}

func (u *UserAddrCtrl) UserAddrList(c *gin.Context) interface{} {
	userId, _ := strconv.Atoi(c.Param("userId"))
	res, err := u.userAddrService.ListByUserId(userId)
	if err != nil {
		return e.ServerError(err.Error())
	}
	return res
}

func (u *UserAddrCtrl) UserAddrAdd(c *gin.Context) interface{} {
	userId, _ := strconv.Atoi(c.Param("userId"))
	uaDto := &dto.UserAddrDto{}
	if err := c.BindJSON(uaDto); err != nil {
		return e.ParameterError(err.Error())
	}
	ua := &models.UserAddress{
		UserId:    userId,
		Consignee: uaDto.Consignee,
		Province:  uaDto.Province,
		City:      uaDto.City,
		District:  uaDto.District,
		Street:    uaDto.Street,
		Zipcode:   uaDto.Zipcode,
		Mobile:    uaDto.Mobile,
		Default:   uaDto.Default,
	}
	err := u.userAddrService.CreateUserAddr(ua)
	if err != nil {
		return e.ServerError(err.Error())
	}
	return "success"
}
