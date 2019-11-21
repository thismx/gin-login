package helpers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func PasswordVerify(inputPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(inputPassword))
	fmt.Printf("%#v", err)
	return err == nil
}

func SuccessMsg(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"msg":     msg,
		"success": true,
	})
}

func ErrorMsg(c *gin.Context, msg string) {
	c.JSON(422, gin.H{
		"msg":     msg,
		"success": false,
	})
}

func SessionSet(c *gin.Context, key interface{}, val interface{}) {
	session := sessions.Default(c)
	session.Set(key, val)
	session.Save()
}

func SessionGet(c *gin.Context, key interface{}) interface{} {
	return sessions.Default(c).Get(key)
}
