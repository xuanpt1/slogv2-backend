package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slogv2/src/main/entity"
	"slogv2/src/main/service"
	"slogv2/src/main/utils/customError"
)

func CreateArticle(c *gin.Context) {
	var article entity.Article
	_ = c.ShouldBind(&article)
	status, err := service.CreateArticle(&article)

	if status == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{
			"code": status,
			"msg":  "文章创建成功",
		})
	} else {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": status,
				"msg":  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": status,
				"msg":  customError.GetMsg(status),
			})
		}
	}
}

func GetArticle(c *gin.Context) {
	param := c.Param("aid")
	query := c.Query("aid")
	var aid string
	if query != "" {
		aid = query
	} else if param != "" {
		aid = param
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": customError.ARTICLE_NOT_FOUND,
			"msg":  "参数错误",
		})
	}
	article, status, err := service.GetArticleByAid(aid)

	if status == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{
			"code": status,
			"msg":  "文章获取成功",
			"data": article,
		})
	} else {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": status,
				"msg":  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": status,
				"msg":  customError.GetMsg(status),
				"data": article,
			})
		}
	}
}
