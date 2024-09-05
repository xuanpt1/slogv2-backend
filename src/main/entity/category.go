package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryId   int    `gorm:"type:int;not null;primaryKey" json:"category_id" label:"分类id"`
	CategoryName string `gorm:"type:varchar(50);not null" json:"category_name" label:"分类名称"`
	CategoryDesc string `gorm:"type:varchar(100)" json:"category_desc" label:"分类描述"`
	//TODO 待定 实现本地文件管理后才有意义
	CategoryIcon string `gorm:"type:varchar(100)" json:"category_icon" label:"分类图标"`
	Count        int    `gorm:"type:int;not null" json:"count" label:"分类下的文章数量"`
	IsActive     bool   `gorm:"type:bool;not null" json:"is_active" label:"是否激活"`

	//default 0 表示无父分类
	ParentId int `gorm:"type:int;not null" json:"parent_id" label:"父分类id"`
}
