package entity

import "online-learning-restful-api/core"

type MasterLikeableType struct {
	core.EntityModel `gorm:"embedded"`
	LikeableName     string `gorm:"varchar(50);not null"`
}

func (MasterLikeableType) TableName() string {
	return "master_likeable_type"
}
