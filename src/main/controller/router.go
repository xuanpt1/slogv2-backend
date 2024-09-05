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

	//文章CRUD
	routerv1.POST("/article", CreateArticle)
	routerv1.GET("/article/:aid", GetArticle)
	routerv1.GET("/article", GetArticle)
	routerv1.PUT("/article", UpdateArticle)
	routerv1.DELETE("/article/:aid", DeleteArticle)
	routerv1.DELETE("/article", DeleteArticle)
	routerv1.POST("/articleList", GetArticleListByPage)
	routerv1.GET("/article/likes/:aid", UpdateArticleLikes)
	routerv1.GET("/article/likes", UpdateArticleLikes)
	routerv1.POST("/articleWithCategory", CreateArticleWithCategory)
	routerv1.PUT("/articleWithCategory", UpdateArticleWithCategory)
	routerv1.GET("/articleWithCategory", GetArticleWithCategory)
	routerv1.GET("/achieve", GetAchieveArticleList)

	//分类CRUD
	routerv1.POST("/category", CreateCategory)
	routerv1.PUT("/category", UpdateCategory)
	routerv1.GET("/category", GetCategory)
	routerv1.GET("/category/:cid", GetCategory)
	routerv1.DELETE("/category", DeleteCategory)
	routerv1.DELETE("/category/:cid", DeleteCategory)
	routerv1.GET("/categoryList", GetCategoryList)

	_ = router.Run(utils.HttpPort)
}
