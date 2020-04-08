package initDB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hackathon/model"
	"log"
)

//DB数据库
var Db *gorm.DB


//初始化数据库
func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:210377091ztc@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err.Error())
	}
	Db.AutoMigrate(&model.User{})
}

//关闭数据库
func CloseDB() {
	Db.Close()
}