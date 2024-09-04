package controller

import (
	"github.com/gin-gonic/gin"
	"slogv2/src/main/utils"
)

func InitRouter() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.RedirectTrailingSlash = false

	routerv1 := router.Group("api/v1")
	routerv1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routerv1.PUT("/article", CreateArticle)
	routerv1.GET("/article/:aid", GetArticle)
	routerv1.GET("/article", GetArticle)

	_ = router.Run(utils.HttpPort)
}
