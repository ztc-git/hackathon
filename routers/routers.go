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
	router.GET("/api/get", func(c *gin.Context) {
		c.JSON(200,gin.H{"msg":"跨域"})
	})

	return router
}
