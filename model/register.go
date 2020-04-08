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
