package model

import "github.com/jinzhu/gorm"

//用户信息表
type User struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(20);not null"`
	Phone string `gorm:"type:varchar(11);not null;unique"`
	Password string `grom:"size:255; not null"`
}