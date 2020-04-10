package controller

import (
	"github.com/gin-gonic/gin"
	"hackathon/initDB"
	"hackathon/jwt"
	"hackathon/model"
	"log"
	"net/http"
)

//登陆
func Login(c *gin.Context) {
	//获取参数
	var userLogin model.UserLogin
	err := c.BindJSON(&userLogin)

	//判断手机号是否存在
	var user model.User
	initDB.Db.Where("user_phone=?", userLogin.Phone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "wname"})
		return
	}

	//判断密码是否正确
	if user.UserPassword != userLogin.Password {
		c.JSON(http.StatusBadRequest, gin.H{"status": "wpswd"})
		return
	}

	//发放token
	token, Err := jwt.ReleaseToken(user)
	if Err != nil {
		log.Printf("token generate erroer : %v", err)
		return
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"statuse": "success",
		"token": token,
		"nickname": user.UserNickname,
		"userid": user.ID,
		"phone":user.UserPhone,
	})
}
