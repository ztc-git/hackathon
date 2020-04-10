package model

import "github.com/jinzhu/gorm"


type VerifyPhone  struct {
	gorm.Model
	Phone string `json:"phone" binding:"required,len=11"`
}

type BindingUser struct {
	Phone string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type UserLogin struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password"`
}

type NewPassword struct {
	NewPassword string `json:"newpassword" binding:"required"`
}

type NicknameSignature struct {
	NickName string `json:"nick_name"`
	Signature string `json:"signature"`
}

type PraisePoints struct {
	//CategoriesOfIslands int `json:"categories_of_islands"`
	ArticleId uint `json:"article_id"`
}

