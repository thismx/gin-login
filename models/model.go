package models

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func InitDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "./test.db")
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	if err != nil {
		panic("err open databases")
	}
	DB = db
	return DB
}
