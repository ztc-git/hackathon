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

