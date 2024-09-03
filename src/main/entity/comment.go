package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Cid      int    `gorm:"type:int;not null;primary key" json:"cid" label:"评论id"`
	Aid      int    `gorm:"type:int;not null" json:"aid" label:"文章id"`
	Content  string `gorm:"type:text;not null" json:"content" label:"评论内容"`
	Uid      int    `gorm:"type:int;not null" json:"uid" label:"用户id"`
	Uname    string `gorm:"type:varchar(20);not null" json:"uname" label:"用户名(昵称)"`
	Mail     string `gorm:"type:varchar(50)" json:"mail" label:"邮箱"`
	Homepage string `gorm:"type:varchar(100)" json:"homepage" label:"个人主页(博客) 可为空"`
	Parent   int    `gorm:"type:int;not null" json:"parent" label:"父评论id"`
	Root     int    `gorm:"type:int;not null" json:"root" label:"根评论id"`
	IsActive bool   `gorm:"type:bool;not null" json:"is_active" label:"是否激活"`
	Agent    string `gorm:"type:varchar(100)" json:"agent" label:"用户代理"`
	Likes    int    `gorm:"type:int;not null" json:"likes" label:"点赞数"`
	Dislikes int    `gorm:"type:int;not null" json:"dislikes" label:"点踩数"`
}
