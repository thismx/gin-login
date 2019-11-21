package routers

import (
	"gintest/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	router.GET("/", controllers.Index)
	router.GET("/welcome", controllers.Welcome)
	router.GET("/register", controllers.RegisterForm)
	router.POST("/login", controllers.Login)
	router.POST("/register_action", controllers.Register)
}
