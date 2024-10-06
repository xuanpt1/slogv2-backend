package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"slogv2/src/main/entity"
	"slogv2/src/main/service"
	"slogv2/src/main/vo"
)

func CreateArticle(c *gin.Context) {
	var article entity.Article
	_ = c.ShouldBind(&article)

	status, err := service.CreateArticle(&article)

	if !ResponseHandler(c, status, err, "文章创建成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetArticle(c *gin.Context) {
	aid := ParamHandler(c, "aid")["aid"]

	article, status, err := service.GetArticleByAid(aid)

	if !ResponseHandler(c, status, err, "文章获取成功", article) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func UpdateArticle(c *gin.Context) {
	var article entity.Article
	_ = c.ShouldBind(&article)

	status, err := service.UpdateArticle(&article)

	if !ResponseHandler(c, status, err, "文章更新成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func DeleteArticle(c *gin.Context) {
	aid := ParamHandler(c, "aid")["aid"]

	status, err := service.DeleteArticle(aid)

	ResponseHandler(c, status, err, "文章删除成功", nil)
}

func UpdateArticleLikes(c *gin.Context) {
	aid := ParamHandler(c, "aid")["aid"]

	status, err := service.UpdateArticleLikes(aid)

	if !ResponseHandler(c, status, err, "文章点赞成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetArticleListByPage(c *gin.Context) {
	var page vo.Page
	_ = c.ShouldBind(&page)

	articleList, _, status, err := service.GetArticleList(page)

	if !ResponseHandler(c, status, err, "文章获取成功", articleList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetAchieveArticleList(c *gin.Context) {
	achieveList, status, err := service.GetAchieveArticleList()

	if !ResponseHandler(c, status, err, "归档信息获取成功", achieveList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func CreateArticleWithCategory(c *gin.Context) {
	var articleWithCategory vo.ArticleWithCategory
	_ = c.ShouldBind(&articleWithCategory)

	status, err := service.CreateArticleWithCategory(&articleWithCategory)

	if !ResponseHandler(c, status, err, "文章创建成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func UpdateArticleWithCategory(c *gin.Context) {
	var articleWithCategory vo.ArticleWithCategory
	_ = c.ShouldBind(&articleWithCategory)

	status, err := service.UpdateArticleWithCategory(&articleWithCategory)

	if !ResponseHandler(c, status, err, "文章更新成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetArticleWithCategory(c *gin.Context) {
	aid := ParamHandler(c, "aid")["aid"]

	articleWithCategory, status, err := service.GetArticleWithCategory(aid)

	if !ResponseHandler(c, status, err, "文章获取成功", articleWithCategory) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetAllArticle(c *gin.Context) {
	articleList, total, status, err := service.GetAllArticle()

	if !ResponseHandler(c, status, err, fmt.Sprintf("共%d篇文章获取成功", total), articleList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetArticleListInfo(c *gin.Context) {

	articleListInfo, status, err := service.GetArticleListInfo()

	if !ResponseHandler(c, status, err, "文章列表信息获取成功", articleListInfo) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}
