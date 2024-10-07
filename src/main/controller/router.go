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

	//for test
	routerv1.GET("/getAllArticle", GetAllArticle)

	//文章CRUD
	routerv1.POST("/article", CreateArticle)
	routerv1.GET("/article/:aid", GetArticle)
	routerv1.GET("/article", GetArticle)
	routerv1.PUT("/article", UpdateArticle)
	routerv1.DELETE("/article/:aid", DeleteArticle)
	routerv1.DELETE("/article", DeleteArticle)
	routerv1.POST("/articleList", GetArticleListByPage)
	routerv1.GET("/articleList", GetArticleListInfo)
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

	//配置CRUD
	routerv1.POST("/options", CreateOption)
	routerv1.PUT("/options", UpdateOption)
	routerv1.GET("/options", GetOptionsByName)
	routerv1.GET("/options/:name", GetOptionsByName)
	routerv1.DELETE("/options", DeleteOption)
	routerv1.DELETE("/options/:name", DeleteOption)
	routerv1.GET("/optionsList", GetOptionsList)

	//评论CRUD
	routerv1.POST("/comment", CreateComment)
	routerv1.PUT("/comment", UpdateComment)
	routerv1.GET("/comment", GetComment)
	routerv1.GET("/comment/:cid", GetComment)
	routerv1.DELETE("/comment", DeleteComment)
	routerv1.DELETE("/comment/:cid", DeleteComment)
	routerv1.GET("/commentList", GetCommentList)
	routerv1.GET("/commentList/:aid", GetCommentListByAid)
	routerv1.GET("/commentTree", GetCommentTreeByAid)
	routerv1.GET("/commentTree/:aid", GetCommentTreeByAid)
	routerv1.GET("/commentPending", GetPendingCommentList)
	routerv1.GET("/commentEnable", EnableComment)
	routerv1.GET("/commentEnable/:cid", EnableComment)
	routerv1.GET("/commentCount", GetCommentCount)
	routerv1.GET("/commentCount/:aid", GetCommentCountByAid)

	//用户CRUD
	routerv1.POST("/user", CreateUser)
	routerv1.PUT("/user", UpdateUser)
	routerv1.GET("/user", GetUserByUid)
	routerv1.GET("/user/:uid", GetUserByUid)
	routerv1.DELETE("/user", DeleteUser)
	routerv1.DELETE("/user/:uid", DeleteUser)
	routerv1.GET("/userList", GetUserList)
	routerv1.GET("/userByName", GetUserByUsername)
	routerv1.GET("/userByName/:username", GetUserByUsername)

	routerv1.POST("/login", Login)

	_ = router.Run(utils.HttpPort)
}
