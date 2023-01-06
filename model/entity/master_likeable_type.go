package entity

import "online-learning-restful-api/core"

type MasterLikeableType struct {
	core.EntityModel
	LikeableName string `gorm:"varchar(50)"`
}

func (MasterLikeableType) TableName() string {
	return "master_likeable_type"
}
