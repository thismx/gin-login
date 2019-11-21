package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context){

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Welcome(c *gin.Context){
	session := sessions.Default(c)
	if(session.Get("user") == nil){
		c.Redirect(302, "/")
	}
	c.HTML(200, "welcome.html", gin.H{
		"username" : session.Get("user"),
	})
}