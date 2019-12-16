package main

import (
	"gintest/helpers"
	"gintest/models"
	"gintest/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"html/template"
)

func main() {
	r := gin.Default()

	//数据库连接
	db := models.InitDB()
	defer db.Close()

	//初初化
	initialization(r)

	r.Run("localhost:8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func initialization(r *gin.Engine) {

	//session
	store := cookie.NewStore([]byte("secr@115156@KFJKS#$@"))
	r.Use(sessions.Sessions("cgn_session_id", store))

	//模板函数
	r.SetFuncMap(template.FuncMap{
		"round": helpers.Round,
	})

	//模板目录
	r.LoadHTMLGlob("./views/**/*")

	//静态目录
	r.Static("/static", "./public/static")

	//路由处理
	routers.SetRouter(r)

}
