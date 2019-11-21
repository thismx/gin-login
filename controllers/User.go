package controllers

import (
	"gintest/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterForm(c *gin.Context){
	c.HTML(200, "welcome.html", gin.H{})
}

func Login(c *gin.Context){
	var (
		err  error
		user *models.User
	)

	username := c.PostForm("email")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(422,  gin.H{
			"msg": "请输入账号和密码",
			"url" : "",
			"data" : "",
		})
		return
	}

	user, err = models.GetUserByUsername(username)

	if err != nil || user.Pass != password {
		c.JSON(422,  gin.H{
			"msg": "请输入账号或密码有误, 请重试",
			"url" : "",
			"data" : "",
		})
		return
	}
	session := sessions.Default(c)
	session.Clear()
	session.Set("user", user.Name)
	session.Save()

	c.JSON(200,  gin.H{
		"msg": "登录成功",
		"url" : "",
		"data" : "",
	})
}