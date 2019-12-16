package controllers

import (
	"gintest/helpers"
	"gintest/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (this *UserController) RegisterForm(c *gin.Context) {
	c.HTML(200, "user/register.html", gin.H{})
}

func (this *UserController) Register(c *gin.Context) {
	username := c.PostForm("email")
	password := c.PostForm("password")
	password2 := c.PostForm("password_confirm")
	if username == "" || password == "" {
		helpers.ErrorMsg(c, "请输入账号和密码")
		return
	}

	if password2 != password {
		helpers.ErrorMsg(c, "确认密码不一致")
		return
	}
	user := models.User{
		Name: username,
		Pass: helpers.PasswordHash(password),
	}

	if err := models.DB.Create(&user).Error; err != nil {
		helpers.ErrorMsg(c, "账号已存在, 请重新注册")
		return
	}

	helpers.SuccessMsg(c, "注册成功")
}

func (this *UserController) LoginForm(c *gin.Context) {
	c.HTML(200, "user/login.html", gin.H{})
}

func (this *UserController) Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}

func (this *UserController) Login(c *gin.Context) {
	var (
		err  error
		user *models.User
	)

	username := c.PostForm("email")
	password := c.PostForm("password")
	if username == "" || password == "" {
		helpers.ErrorMsg(c, "请输入账号和密码")
		return
	}

	user, err = models.GetUserByUsername(username)
	if err != nil || helpers.PasswordVerify(password, user.Pass) == false {
		helpers.ErrorMsg(c, "请输入账号或密码有误, 请重试")
		return
	}

	helpers.SessionSet(c, "user", user.Name)

	helpers.SuccessMsg(c, "登录成功")
}
