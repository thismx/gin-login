package main

import (
	"gintest/models"
	"gintest/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := gin.Default()
	initialization(r)
	r.Run("192.168.100.120:9090") // 监听并在 0.0.0.0:8080 上启动服务
}

func initialization(r *gin.Engine){

	_, err := models.InitDB()
	if err != nil {
		panic("err open databases")
		return
	}
	//defer db.Close()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("cgn_session_id", store))

	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./public/static")

	routers.SetRouter(r)
}