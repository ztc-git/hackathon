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
	Db.AutoMigrate(&model.UserFriend{})
	Db.AutoMigrate(&model.Comment{})
	Db.AutoMigrate(&model.SecretIsland{})
	Db.AutoMigrate(&model.StoryIsland{})
	Db.AutoMigrate(&model.StorySolitaire{})
	Db.Model(&model.UserFriend{}).AddForeignKey("friend_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.Model(&model.UserFriend{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.Model(&model.SecretIsland{}).AddForeignKey("author_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.Model(&model.Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.Model(&model.Comment{}).AddForeignKey("story_id", "secret_islands(id)", "RESTRICT", "RESTRICT")
	Db.Model(&model.StorySolitaire{}).AddForeignKey("story_island_id", "story_islands(id)","RESTRICT", "RESTRICT")
}

//关闭数据库
func CloseDB() {
	Db.Close()
}