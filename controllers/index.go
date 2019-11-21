package controllers

import (
	"gintest/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Welcome(c *gin.Context) {

	if helpers.SessionGet(c, "user") == nil {
		c.Redirect(302, "/")
	}
	c.HTML(200, "welcome.html", gin.H{
		"username": helpers.SessionGet(c, "user"),
	})
}
