package entity

import "gorm.io/gorm"

type Options struct {
	gorm.Model
	Key   string `gorm:"type:varchar(50);not null" json:"key" label:"键"`
	Value string `gorm:"type:varchar(100);not null" json:"value" label:"值"`
	Uid   int    `gorm:"type:int;not null" json:"uid" label:"用户id"`
}
