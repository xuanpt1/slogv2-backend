package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//TIP 使用gorm.Model 自带主键ID
	//仍需使用uid 便于与其他表的关联
	Uid      string `gorm:"type:varchar(20);primaryKey;not null" json:"uid" label:"用户uid"`
	Username string `gorm:"type:varchar(20);not null;unique" json:"username" label:"用户名 不允许重复 登陆用"`
	Password string `gorm:"type:varchar(100);not null" json:"password" label:"密码"`
	Nickname string `gorm:"type:varchar(20);not null" json:"nickname" label:"昵称 显示的名称 默认是用户名"`
	//TODO 待定 预留给多用户情形使用 暂时不用
	Url      string `gorm:"type:varchar(100)" json:"url" label:"个人主页"`
	Email    string `gorm:"type:varchar(50)" json:"email" label:"邮箱"`
	Avatar   string `gorm:"type:varchar(100)" json:"avatar" label:"头像"`
	IsAdmin  bool   `gorm:"type:bool;not null" json:"is_admin" label:"是否管理员"`
	IsActive bool   `gorm:"type:bool;not null" json:"is_active" label:"是否激活"`
}
