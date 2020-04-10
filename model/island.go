package model

import "github.com/jinzhu/gorm"

//分享岛
type  SharingIsland struct {
	gorm.Model
	TopicOfConversation	string 	`json:"topic_of_conversation"`
	ImagesAddress string
	AuthorID uint `gorm:"not null"`
	Title string `gorm:"type:varchar(30); not null" json:"title"`
	Text string `gorm:"type:text; not null" json:"text"`
	PraisePoints uint `gorm:"type:int; default:0"`
	Collection uint `gorm:"type:int; default:0"`
}

//秘密之岛
type SecretIsland struct {
	gorm.Model
	TopicOfConversation	string 	`json:"topic_of_conversation"`
	ImagesAddress string 		`json:"images_address"`
	AuthorID uint 				`gorm:"not null"`
	Title string 				`gorm:"type:varchar(30); not null" json:"title"`
	Text string 				`gorm:"type:text; not null" json:"text"`
	PraisePoints uint 			`gorm:"type:int; default:0"`
	Collection uint 			`gorm:"type:int; default:0"`
}

//故事岛
type StoryIsland struct {
	gorm.Model
	AuthorID uint `gorm:"not null"`
	Title string `gorm:"type varchar(30); not null" json:"title"`
	StartingParagraph string `gorm:"type:text; not null" json:"starting_paragraph"`
	PraisePoints uint `gorm:"type:int; default:0"`
	Collection uint `gorm:"type:int; default:0"`
}

type StorySolitaire struct {
	gorm.Model
	StoryIslandID uint
	Text string `gorm:"not null"`
	PraisePoints uint `gorm:"type:int; default:0"`
	Collection uint `gorm:"type:int; default:0"`
}