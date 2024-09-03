package entity

type Relationship struct {
	Cid int `gorm:"type:int;not null;primary key;autoIncrement:false" json:"cid" label:"评论id"`
	Aid int `gorm:"type:int;not null;primary key;autoIncrement:false" json:"aid" label:"文章id"`
}
