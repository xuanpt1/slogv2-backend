package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"slogv2/src/main/entity"
	"slogv2/src/main/service"
)

func GetComment(c *gin.Context) {
	cid := ParamHandler(c, "cid")["cid"]

	comment, status, err := service.GetCommentByCid(cid)

	if !ResponseHandler(c, status, err, "评论获取成功", comment) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func UpdateComment(c *gin.Context) {
	var comment entity.Comment
	_ = c.ShouldBind(&comment)

	status, err := service.UpdateComment(&comment)

	if !ResponseHandler(c, status, err, "评论更新成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func DeleteComment(c *gin.Context) {
	cid := ParamHandler(c, "cid")["cid"]

	status, err := service.DeleteComment(cid)

	if !ResponseHandler(c, status, err, "评论删除成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func CreateComment(c *gin.Context) {
	var comment entity.Comment
	_ = c.ShouldBind(&comment)

	status, err := service.CreateComment(&comment)

	if !ResponseHandler(c, status, err, "评论创建成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetCommentList(c *gin.Context) {
	commentList, status, err := service.GetCommentList()

	if !ResponseHandler(c, status, err, "评论列表获取成功", commentList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetCommentListByAid(c *gin.Context) {
	aid := ParamHandler(c, "aid")["aid"]

	commentList, status, err := service.GetCommentListByAid(aid)

	if !ResponseHandler(c, status, err, "评论列表获取成功", commentList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetCommentCountByAid(c *gin.Context) {
	aid := ParamHandler(c, "aid")["aid"]

	commentCount, status, err := service.GetCommentCountByAid(aid)

	if !ResponseHandler(c, status, err, "评论数量获取成功", commentCount) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetCommentCount(c *gin.Context) {
	commentCount, status, err := service.GetCommentCount()

	if !ResponseHandler(c, status, err, "评论数量获取成功", commentCount) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetCommentTreeByAid(c *gin.Context) {
	aid := ParamHandler(c, "aid")["aid"]

	commentTree, status, err := service.GetCommentTreeByAid(aid)

	if !ResponseHandler(c, status, err, "评论树获取成功", commentTree) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetPendingCommentList(c *gin.Context) {
	commentList, status, err := service.GetPendingCommentList()

	if !ResponseHandler(c, status, err, "待审核评论列表获取成功", commentList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func EnableComment(c *gin.Context) {
	cid := ParamHandler(c, "cid")["cid"]

	status, err := service.EnableComment(cid)

	if !ResponseHandler(c, status, err, "评论审核成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}
