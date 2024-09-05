package vo

import "slogv2/src/main/entity"

type AchieveArticle struct {
	Year int `json:"year" label:"年份"`
	//Month       int              `json:"month" label:"月份"`
	ArticleList []entity.Article `json:"article_list" label:"文章列表"`
}
