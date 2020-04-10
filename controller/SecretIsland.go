package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hackathon/initDB"
	"hackathon/model"
	"hackathon/util"
	"log"
	"net/http"
)

//秘密岛发表秘密
func ReleaseSecret(c *gin.Context) {
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

//获取所有内容
func GetPosts(c *gin.Context) {
	user := util.GetUser(c)

	var articleList []model.SecretIsland
	initDB.Db.Where("author_id = ?", user.ID).Find(&articleList)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data": articleList,
	})
}

func GetComments(c *gin.Context) {
	var comment model.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var commentList []model.Comment
	initDB.Db.Where("story_id = ?", comment.StoryID).Find(&commentList)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":commentList,
	})
}
