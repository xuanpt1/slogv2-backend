package vo

import "slogv2/src/main/entity"

type ArticleWithCategory struct {
	Article      entity.Article    `json:"article"`
	CategoryList []entity.Category `json:"category"`
}
