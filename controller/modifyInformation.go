package controller

import (
	"github.com/gin-gonic/gin"
	"hackathon/initDB"
	"hackathon/jwt"
	"hackathon/util"
	"log"
)


type NewPassword struct {
	NewPassword string `json:"newpassword" binding:"required"`
}
//修改密码
func ChangePassword(c *gin.Context) {
	//得到user
	user := util.GetUser(c)

	//得到新密码
	var newPassword NewPassword
	err := c.BindJSON(&newPassword)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//提交修改
	initDB.Db.Model(&user).Update("user_password", newPassword.NewPassword)

	//生成新token
	token, Err := jwt.ReleaseToken(user)
	if Err != nil {
		log.Printf("token generate erroer : %v", err)
		return
	}

	//返回结果
	c.JSON(200, gin.H{
		"Msg":"密码修改成功",
		"NewToken": token,
		})
}


type NicknameSignature struct {
	NickName string `json:"nick_name"`
	Signature string `json:"signature"`
}
//修改个性签名或昵称
func ChangeNicknameOrSignature(c *gin.Context) {
	//得到user
	user := util.GetUser(c)

	var nicknameSignature NicknameSignature
	err := c.BindJSON(&nicknameSignature)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//修改昵称
	if nicknameSignature.NickName != "" {
		initDB.Db.Model(&user).Update("user_nickname", nicknameSignature.NickName)
	}

	//修改签名
	if nicknameSignature.Signature != "" {
		initDB.Db.Model(&user).Update("user_personal_signature", nicknameSignature.Signature)
	}
	c.JSON(200, gin.H{"Msg": "修改信息成功"})
}