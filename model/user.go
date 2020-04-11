package model

import "github.com/jinzhu/gorm"

//好友关关联表
type UserFriend struct {
	ID uint `gorm:"primary key;"`
	UserID uint	 `gorm:"unique; not null"`
	FriendID uint	`gorm:"unique; not null"`
	FriendNote string `gorm:"type:varchar(20)"`
}


//评论
type Comment struct {
	gorm.Model
	UserID uint `gorm:"not null" json:"user_id"`
	StoryID uint `gorm:"not null" json:"story_id"`
	FatherCommentID uint `gorm:"not null; default:0"json:"father_comment_id"`
	PraisePoints uint `gorm:"default:0"`
	CommentText string `gorm:"type:text; not null" json:"comment_text"`
}


//用户信息表
type User struct{
	gorm.Model
	UserPhone string `gorm:"type:varchar(11); not null ; unique" json:"phone" form:"phone" binding:"required,len=11"`
	UserPassword string `gorm:"type: varchar(20); not null " form:"password" json:"password" binding:"required"`
	UserNickname string `gorm:"type: varchar(20); not null" form:"nickname" json:"nickname"`
	UserPersonalSignature string `gorm:"type:varchar(50)"`
	UserImage string
}
