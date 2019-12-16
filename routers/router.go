package routers

import (
	"gintest/controllers"
	"gintest/helpers"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	UserController := new(controllers.UserController)
	IndexController := new(controllers.IndexController)

	router.GET("/", IndexController.Index)
	router.GET("/register", UserController.RegisterForm)
	router.GET("/login", UserController.LoginForm)
	router.GET("/logout", UserController.Logout)
	router.POST("/login_action", UserController.Login)
	router.POST("/register_action", UserController.Register)

	authorized := router.Group("/")
	authorized.Use(AuthRequired())
	{
		router.GET("/welcome", IndexController.Welcome)
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if helpers.SessionGet(c, "user") != nil {
			c.Next()
			return
		}
		c.Redirect(302, "/login")
	}
}
