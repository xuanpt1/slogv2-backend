package test

import (
	"github.com/goccy/go-json"
	"slogv2/src/main/entity"
	"slogv2/src/main/utils"
)

func GenTestArticle() string {
	article := entity.Article{
		Title:        "测试文章",
		Content:      "测试文章内容",
		Uid:          1,
		Likes:        0,
		Clicks:       0,
		AllowComment: true,
		Abstract:     "测试文章摘要",
		Image:        "URL_ADDRESS",
	}
	bytes, err := json.Marshal(article)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func RegexTest(source string) string {
	return utils.GetPureTextRegex(source)
}
