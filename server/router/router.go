package router

import (
	"server/internal/user"
	"server/internal/ws"

	"github.com/gin-gonic/gin"
)

var routerInstance *gin.Engine

func InitRouter(userController *user.Controller, wsController *ws.Controller) {
	routerInstance = gin.Default()
	routerInstance.Use(CORSMiddleware())

	routerInstance.POST("/signup", userController.CreateUser)
	routerInstance.POST("/login", userController.Login)
	routerInstance.GET("/logout", userController.Logout)

	routerInstance.GET("/chats", userController.ListUsers)
	routerInstance.GET("/ws/joinChat/:chatId", wsController.JoinRoom)
	routerInstance.GET("chat/:chatId", wsController.GetMessages)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RunRouter(addr string) error {
	return routerInstance.Run(addr)
}
