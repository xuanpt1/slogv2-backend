package customError

import (
	"errors"
	"fmt"
)

const (
	//通用错误码
	SUCCESS = 200
	FAIL    = 500

	//文章错误码
	ARTICLE_NOT_FOUND    = 1001
	ARTICLE_UPDATE_FAIL  = 1002
	ARTICLE_DELETE_FAIL  = 1003
	ARTICLE_CREATE_FAIL  = 1004
	ARTICLE_LIST_FAIL    = 1005
	ARTICLE_ARCHIVE_FAIL = 1006
	ARTICLE_LIKE_FAIL    = 1007

	//分类错误码
	CATEGORY_NOT_FOUND   = 2001
	CATEGORY_UPDATE_FAIL = 2002
	CATEGORY_DELETE_FAIL = 2003
	CATEGORY_CREATE_FAIL = 2004
	CATEGORY_LIST_FAIL   = 2005
	CATEGORY_ADD_FAIL    = 2006
	CATEGORY_REMOVE_FAIL = 2007

	//文章-分类关系错误码
	RELATION_NOT_FOUND   = 3001
	RELATION_CREATE_FAIL = 3002
	RELATION_DELETE_FAIL = 3003

	//评论错误码
	COMMENT_NOT_FOUND   = 4001
	COMMENT_UPDATE_FAIL = 4002
	COMMENT_DELETE_FAIL = 4003
	COMMENT_CREATE_FAIL = 4004
	COMMENT_LIST_FAIL   = 4005

	//配置错误码
	OPTIONS_NOT_FOUND   = 5001
	OPTIONS_UPDATE_FAIL = 5002
	OPTIONS_DELETE_FAIL = 5003
	OPTIONS_CREATE_FAIL = 5004
	OPTIONS_LIST_FAIL   = 5005

	//用户错误码
	USER_NOT_FOUND   = 6001
	USER_UPDATE_FAIL = 6002
	USER_DELETE_FAIL = 6003
	USER_CREATE_FAIL = 6004
	USER_LIST_FAIL   = 6005

	//其他错误码
	OTHER_ERROR = 114514
)

var codeMsg = map[int]string{
	SUCCESS: "success",
	FAIL:    "fail",

	ARTICLE_NOT_FOUND:    "文章不存在",
	ARTICLE_UPDATE_FAIL:  "文章更新失败",
	ARTICLE_DELETE_FAIL:  "文章删除失败",
	ARTICLE_CREATE_FAIL:  "文章添加失败",
	ARTICLE_LIST_FAIL:    "文章列表获取失败",
	ARTICLE_ARCHIVE_FAIL: "文章归档获取失败",
	ARTICLE_LIKE_FAIL:    "文章点赞失败",

	CATEGORY_NOT_FOUND:   "分类不存在",
	CATEGORY_UPDATE_FAIL: "分类更新失败",
	CATEGORY_DELETE_FAIL: "分类删除失败",
	CATEGORY_CREATE_FAIL: "分类创建失败",
	CATEGORY_LIST_FAIL:   "分类列表获取失败",
	CATEGORY_ADD_FAIL:    "添加分类失败",
	CATEGORY_REMOVE_FAIL: "移除分类失败",

	RELATION_NOT_FOUND:   "无相关分类信息",
	RELATION_CREATE_FAIL: "文章添加分类失败失败",
	RELATION_DELETE_FAIL: "文章移除分类失败失败",

	COMMENT_NOT_FOUND:   "评论不存在",
	COMMENT_UPDATE_FAIL: "评论更新失败",
	COMMENT_DELETE_FAIL: "评论删除失败",
	COMMENT_CREATE_FAIL: "评论添加失败",
	COMMENT_LIST_FAIL:   "评论列表获取失败",

	OPTIONS_NOT_FOUND:   "配置不存在",
	OPTIONS_UPDATE_FAIL: "配置更新失败",
	OPTIONS_DELETE_FAIL: "配置删除失败",
	OPTIONS_CREATE_FAIL: "配置创建失败",
	OPTIONS_LIST_FAIL:   "配置列表获取失败",

	USER_NOT_FOUND:   "用户不存在",
	USER_UPDATE_FAIL: "用户更新失败",
	USER_DELETE_FAIL: "用户删除失败",
	USER_CREATE_FAIL: "用户创建失败",
	USER_LIST_FAIL:   "用户列表获取失败",

	OTHER_ERROR: "其他错误",
}

func GetMsg(code int) string {
	return codeMsg[code]
}

// GetError 传入额外的错误信息
// param code 自定义错误码
// param msg 额外的错误信息
// return error 自定义错误，错误信息为 "自定义错误码: 自定义错误信息\n 额外的错误信息"
//
//	func GetError(code int) error {
//		return error.New(fmt.Sprintf("%d: %s", code, codeMsg[code]))
//	}
func GetError(code int, msg string) error {
	return errors.New(fmt.Sprintf("%d: %s\n %s", code, codeMsg[code], msg))
}
