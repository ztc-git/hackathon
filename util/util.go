package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hackathon/initDB"
	"hackathon/model"
	"log"
	"math/rand"
	"strings"
	"time"
)


//判断手机号是否已经被注册
func IsTelephoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}


//生成验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
		if err != nil {
			log.Println(err.Error())
		}
	}
	return sb.String()
}

//获取User
func GetUser(c *gin.Context) model.User {
	//得到user_id
	userId,_ := c.Get("user_id")
	userID, _:= userId.(uint)

	var user model.User
	initDB.Db.First(&user, userID)

	return user
}