package util

import (
	"hackathon/initDB"
	"hackathon/model"
)

func IsFatherCommentExist(ID uint) bool {
	var comment model.Comment
	initDB.Db.First(&comment, ID)

	if comment.ID != 0 || ID == 0{
		return true
	}
	return false
}

//返回的结果
//func ReturnInformation(articleList []model.SecretIsland) {
//
//}