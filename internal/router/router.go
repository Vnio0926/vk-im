package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vk-im/internal/controller"
)

//路由有提供一个分组的概念，就好比 user 用户是一个路由组， role 角色是一个路由组

func Routers() *gin.Engine {

	r := gin.Default()
	r.Use(Cors())
	r.Use(Recovery)

	user := r.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.POST("/modifyUserInfo", controller.ModifyUserInfo)
		user.GET("/getUserInfo", controller.GetUserInfo)
		user.GET("/getUserFriends", controller.GetUserFriends)
	}

	return r
}

// Cors 跨域中间件
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		origin := c.Request.Header.Get("Origin")
//		//处理跨域请求：
//		if origin != "" {
//			c.Header("Access-Control-Allow-Origin", "*") // 允许所有域名访问
//			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
//			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
//			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
//			c.Header("Access-Control-Allow-Credentials", "true")
//		}
//		if method == "OPTIONS" {
//			c.JSON(http.StatusOK, "ok")
//		}
//		// 处理panic
//		defer func() {
//			if err := recover(); err != nil {
//				c.JSON(http.StatusOK, gin.H{
//					"code": http.StatusBadRequest,
//					"msg":  "服务器错误",
//				})
//			}
//		}()
//		c.Next() // 继续请求处理/下一个中间件
//	}
//}

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "服务器错误",
			})
		}
	}()
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				//log.Logger.Error("HttpError", zap.Any("HttpError", err))
				log.Println(err)
			}
		}()

		c.Next()
	}
}
