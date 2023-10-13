package routes

import (
	"fantasy/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	vuUser := router.Group("usq")
	{
		vuUser.GET("login", handler.Login)
		vuUser.POST("loginvalidate", handler.LoginValidate)
		vuUser.GET("generate", handler.UsqHandler)
		vuUser.GET("error", handler.UsqHandlerErr)
		vuUser.GET("glogin", handler.Glogin)
		vuUser.GET("alogin", handler.Applelogin)
		vuUser.GET("getcookie", handler.GetCookieHandler)
	}
}
