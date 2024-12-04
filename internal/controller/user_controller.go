package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vk-im/internal/model"
	"vk-im/internal/service"
	"vk-im/util/common/response"
)

// 用户相关的控制层，登录，注册

// TODO 注册
func Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("请求参数错误"))
		return
	}

	err = service.UserService.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(user))
}

// TODO 登录
func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return
	}

	if ok, _ := service.UserService.Login(&user); ok {
		c.JSON(200, response.Success(user))
	}
	c.JSON(http.StatusOK, response.Fail("登陆失败"))
}

func ModifyUserInfo(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return
	}

	if ok, _ := service.UserService.ModifyUserInfo(&user); ok {
		c.JSON(200, response.Success(user))
	}
	c.JSON(http.StatusOK, response.Fail("修改用户信息失败"))
}

func GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	user, ok := service.UserService.GetUserInfo(id)
	if ok == false {
		c.JSON(http.StatusOK, response.FailCodeMsg(400, "用户不存在"))
	}

	c.JSON(http.StatusOK, response.Success(user))
}

func GetUserFriends(c *gin.Context) {
	id := c.Param("user_id")
	friends, ok := service.UserService.GetUserFriends(id)
	if ok == false {
		c.JSON(http.StatusOK, response.FailCodeMsg(400, "用户不存在"))
	}
	c.JSON(http.StatusOK, response.Success(friends))
}
