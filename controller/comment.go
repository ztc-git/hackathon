package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hackathon/initDB"
	"hackathon/model"
	"hackathon/util"
	"log"
)

//发布评论
func SubmitComments(c *gin.Context) {
	var comment model.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if ! util.IsFatherCommentExist(comment.FatherCommentID) {
		log.Println("评论错误")
		return
	}

	user := util.GetUser(c)

	submitComments := model.Comment{
		Model:           gorm.Model{},
		UserID:          user.ID,
		StoryID:         comment.StoryID,
		FatherCommentID: comment.FatherCommentID,
		PraisePoints:    0,
		CommentText:     comment.CommentText,
	}
	initDB.Db.Create(&submitComments)

	c.JSON(200, gin.H{"Msg": "评论成功"})
}


//给评论点赞
func CommentPraisePoint(c *gin.Context) {
	//获取commentID
	var praisePoints model.PraisePoints
	err := c.BindJSON(&praisePoints)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var comment model.Comment
	initDB.Db.First(&comment, praisePoints.CommentId)
	if comment.ID != 0 {
		comment.PraisePoints += 1
	}
	initDB.Db.Model(&comment).Update("praise_points", comment.PraisePoints)

}