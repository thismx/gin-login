package controllers

import (
	"gintest/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (this *IndexController) Index(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/login")
}

func (this *IndexController) Welcome(c *gin.Context) {

	if helpers.SessionGet(c, "user") == nil {
		c.Redirect(302, "/login")
	}
	c.HTML(200, "index/welcome.html", gin.H{
		"username": helpers.SessionGet(c, "user"),
	})
}
