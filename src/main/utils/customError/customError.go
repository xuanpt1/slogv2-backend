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
