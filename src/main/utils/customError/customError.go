package customError

const (
	//通用错误码
	SUCCESS = 200
	FAIL    = 500

	//文章错误码
	ARTICLE_NOT_FOUND    = 1001
	ARTICLE_UPDATE_FAIL  = 1002
	ARTICLE_DELETE_FAIL  = 1003
	ARTICLE_ADD_FAIL     = 1004
	ARTICLE_LIST_FAIL    = 1005
	ARTICLE_ARCHIVE_FAIL = 1006
	ARTICLE_LIKE_FAIL    = 1007
)

var codeMsg = map[int]string{
	SUCCESS: "success",
	FAIL:    "fail",

	ARTICLE_NOT_FOUND:    "文章不存在",
	ARTICLE_UPDATE_FAIL:  "文章更新失败",
	ARTICLE_DELETE_FAIL:  "文章删除失败",
	ARTICLE_ADD_FAIL:     "文章添加失败",
	ARTICLE_LIST_FAIL:    "文章列表获取失败",
	ARTICLE_ARCHIVE_FAIL: "文章归档获取失败",
	ARTICLE_LIKE_FAIL:    "文章点赞失败",
}
