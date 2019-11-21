package models

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open("sqlite3", "\\\\192.168.100.120\\dbdb\\test.db")
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	if err == nil {
		DB = db
		//db.LogMode(true)
		return db, err
	}
	return nil, err
}
