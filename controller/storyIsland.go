package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hackathon/initDB"
	"hackathon/model"
	"log"
)


func CommitStory(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID, _ := userId.(uint)

	var story model.StoryIsland
	err := c.BindJSON(&story)
	if err != nil {
		log.Println(err.Error())
		return
	}

	newStory := model.StoryIsland{
		Model:             gorm.Model{},
		StoryCategory:     story.StoryCategory,
		AuthorID:          userID,
		Title:             story.Title,
		StartingParagraph: story.StartingParagraph,
		PraisePoints:      0,
		Collection:        0,
	}
	initDB.Db.Create(&newStory)

	c.JSON(200, gin.H{"Msg":"发表成功"})
}

func AddStorySolitaire(c *gin.Context) {
	var story model.StorySolitaire
	err := c.BindJSON(&story)
	if err != nil {
		log.Println(err.Error())
		return
	}

	newStorySolitaire := model.StorySolitaire{
		Model:         gorm.Model{},
		StoryIslandID: story.StoryIslandID,
		Text:          story.Text,
		PraisePoints:  0,
		Collection:    0,
	}
	initDB.Db.Create(&newStorySolitaire)

	c.JSON(200, gin.H{"Msg":"发表成功"})
}