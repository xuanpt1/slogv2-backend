package controller

import (
	"github.com/gin-gonic/gin"
	"slogv2/src/main/middleware"
	"slogv2/src/main/utils"
)

func InitRouter() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.RedirectTrailingSlash = false
	router.Use(middleware.CorsMiddleware())

	normalRouter := router.Group("api/v1")
	{
		normalRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		//for test
		normalRouter.GET("/getAllArticle", GetAllArticle)

		//文章 公开权限
		normalRouter.GET("/article/:aid", GetArticle)
		normalRouter.GET("/article", GetArticle)
		normalRouter.POST("/articleList", GetArticleListByPage)
		normalRouter.GET("/articleList", GetArticleListInfo)
		normalRouter.GET("/article/likes/:aid", UpdateArticleLikes)
		normalRouter.GET("/article/likes", UpdateArticleLikes)
		normalRouter.GET("/articleWithCategory", GetArticleWithCategory)
		normalRouter.GET("/achieve", GetAchieveArticleList)

		//分类
		normalRouter.GET("/category", GetCategory)
		normalRouter.GET("/category/:cid", GetCategory)
		normalRouter.GET("/categoryList", GetCategoryList)

		//配置
		normalRouter.GET("/options", GetOptionsByName)
		normalRouter.GET("/options/:name", GetOptionsByName)

		//评论
		normalRouter.POST("/comment", CreateComment)
		normalRouter.GET("/comment", GetComment)
		normalRouter.GET("/comment/:cid", GetComment)
		normalRouter.GET("/commentList", GetCommentList)
		normalRouter.GET("/commentList/:aid", GetCommentListByAid)
		normalRouter.GET("/commentTree", GetCommentTreeByAid)
		normalRouter.GET("/commentTree/:aid", GetCommentTreeByAid)
		normalRouter.GET("/commentCount", GetCommentCount)
		normalRouter.GET("/commentCount/:aid", GetCommentCountByAid)

		//用户
		normalRouter.POST("/login", Login)

		//TODO 细化用户权限校验 注册的用户不能都为管理员
		normalRouter.POST("/user", CreateUser)
	}

	//采用Restful风格接口，故未为认证权限组额外添加路径前缀
	//authRouter := router.Group("api/v1/" + Admin等标识性字段)
	authRouter := router.Group("api/v1")
	authRouter.Use(middleware.JwtAuthMiddleware())
	{
		//文章CRUD
		authRouter.POST("/article", CreateArticle)
		authRouter.PUT("/article", UpdateArticle)
		authRouter.DELETE("/article/:aid", DeleteArticle)
		authRouter.DELETE("/article", DeleteArticle)
		authRouter.POST("/articleWithCategory", CreateArticleWithCategory)
		authRouter.PUT("/articleWithCategory", UpdateArticleWithCategory)

		//分类CRUD
		authRouter.POST("/category", CreateCategory)
		authRouter.PUT("/category", UpdateCategory)
		authRouter.DELETE("/category", DeleteCategory)
		authRouter.DELETE("/category/:cid", DeleteCategory)

		//配置CRUD
		authRouter.POST("/options", CreateOption)
		authRouter.PUT("/options", UpdateOption)
		authRouter.DELETE("/options", DeleteOption)
		authRouter.DELETE("/options/:name", DeleteOption)
		authRouter.GET("/optionsList", GetOptionsList)

		//评论CRUD
		authRouter.PUT("/comment", UpdateComment)
		authRouter.DELETE("/comment", DeleteComment)
		authRouter.DELETE("/comment/:cid", DeleteComment)
		authRouter.GET("/commentPending", GetPendingCommentList)
		authRouter.POST("/commentEnable", EnableComment)
		authRouter.POST("/commentEnable/:cid", EnableComment)

		//用户CRUD
		authRouter.PUT("/user", UpdateUser)
		authRouter.GET("/user", GetUserByUid)
		authRouter.GET("/user/:uid", GetUserByUid)
		authRouter.DELETE("/user", DeleteUser)
		authRouter.DELETE("/user/:uid", DeleteUser)
		authRouter.GET("/userList", GetUserList)
		authRouter.GET("/userByName", GetUserByUsername)
		authRouter.GET("/userByName/:username", GetUserByUsername)
	}

	_ = router.Run(utils.HttpPort)
}
