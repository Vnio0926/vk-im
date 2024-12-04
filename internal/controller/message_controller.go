package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vk-im/internal/service"
	"vk-im/util/common/request"
	"vk-im/util/common/response"
)

// TODO： 获取消息
func GetMessage(c *gin.Context) {
	var messageRequest request.MessageRequest
	err := c.BindQuery(&messageRequest)
	if err != nil {
		log.Println(err)
	}
	message, err := service.MessageService.GetMessage(messageRequest)
	if err != nil {
		c.JSON(http.StatusOK, response.Fail("获取消息失败"))
		return
	}
	c.JSON(http.StatusOK, response.Success(message))

}
