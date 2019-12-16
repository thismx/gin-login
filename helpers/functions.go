package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math"
)

func PasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func PasswordVerify(inputPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(inputPassword))
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

func Sha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256File(file io.Reader) string {
	h := sha256.New()
	_, erro := io.Copy(h, file)
	if erro != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func Round(value float64, n int) float64 {
	if value == 0 {
		return value
	}
	n10 := math.Pow10(n)
	return math.Trunc((value+0.5/n10)*n10) / n10
}
