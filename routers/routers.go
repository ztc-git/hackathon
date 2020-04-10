package routers

import (
	"github.com/gin-gonic/gin"
	"hackathon/controller"
	"hackathon/middleware"
)


func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	router.POST("/api/regist", controller.VerifyPhone)
	router.POST("/api/regist_confirm", controller.Register)
	router.POST("/api/login", controller.Login)


	r := router.Group("/api", middleware.AuthMiddleware())
	{
		r.POST("/release_secret", controller.ReleaseSecret)
		//r.GET("/article", controller.Article)
		r.PUT("/change_password", controller.ChangePassword)
		r.PUT("/change_nickname_signature", controller.ChangeNicknameOrSignature)
		r.PUT("/praise_points", controller.PraisePoints)
		r.POST("/submit_comments", controller.SubmitComments)
		r.PUT("/comment/praise_points", controller.CommentPraisePoint)
		r.GET("/getPost",controller.GetPosts)
		r.GET("/getComments", controller.GetComments)
	}

	//SharingIsland
	{
		r.POST("/share_story", controller.ReleaseSharingStory)
	}
	//storyIsland
	{
		r.POST("/")
	}
	return router
}
