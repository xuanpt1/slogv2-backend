package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Aid          int    `gorm:"type:int;not null;primaryKey" json:"aid" label:"文章id"`
	Title        string `gorm:"type:varchar(100);not null" json:"title" label:"标题"`
	Image        string `gorm:"type:varchar(100);not null" json:"image" label:"文章头图"`
	Abstract     string `gorm:"type:varchar(100);not null" json:"abstract" label:"摘要"`
	Content      string `gorm:"type:text;not null" json:"content" label:"内容"`
	Uid          int    `gorm:"type:int;not null" json:"uid" label:"用户id"`
	Likes        int    `gorm:"type:int;not null" json:"likes" label:"点赞数"`
	Clicks       int    `gorm:"type:int;not null" json:"clicks" label:"点击数"`
	AllowComment bool   `gorm:"type:bool;not null" json:"allow_comment" label:"是否允许评论"`
}
