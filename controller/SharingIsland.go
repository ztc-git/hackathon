package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hackathon/initDB"
	"hackathon/model"
	"log"
)

func ReleaseSharingStory(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID, _ := userId.(uint)

	var story model.SharingIsland
	err := c.BindJSON(&story)
	if err != nil {
		log.Println(err.Error())
		return
	}

	newStory := model.SharingIsland{
		Model:               gorm.Model{},
		TopicOfConversation: story.TopicOfConversation,
		ImagesAddress:       story.ImagesAddress,
		AuthorID:            userID,
		Title:               story.Title,
		Text:                story.Text,
		PraisePoints:        0,
		Collection:          0,
	}

	initDB.Db.Create(&newStory)

	//返回结果
	c.JSON(200, gin.H{"Msg": "故事发表成功"})
}

