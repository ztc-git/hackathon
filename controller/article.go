package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hackathon/initDB"
	"hackathon/model"
	"log"
)

//秘密岛发表秘密
func ReleaseStory(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userID, _ := userId.(uint)

	var story model.SecretIsland
	err := c.BindJSON(&story)
	if err != nil {
		log.Println(err.Error())
		return
	}

	newStory := model.SecretIsland{
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


//点赞
func PraisePoints(c *gin.Context) {
	var praisePoints model.PraisePoints
	err := c.BindJSON(&praisePoints)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var article model.SecretIsland
	initDB.Db.First(&article, praisePoints.ArticleId)
	if article.ID != 0 {
		article.PraisePoints += 1
	}
	initDB.Db.Model(&article).Update("praise_points", article.PraisePoints)

}

