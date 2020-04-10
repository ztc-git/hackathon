package model

import "github.com/jinzhu/gorm"

//用户信息表
type User struct{
	gorm.Model
	UserPhone string `gorm:"type:varchar(11); not null ; unique" json:"phone" form:"phone" binding:"required,len=11"`
	UserPassword string `gorm:"type: varchar(20); not null " form:"password" json:"password" binding:"required"`
	UserNickname string `gorm:"type: varchar(20); not null" form:"nickname" json:"nickname"`
	UserPersonalSignature string `gorm:"type:varchar(50)"`
	UserImage string `gorm:"default:'默认就是默认，别问'"`
}
