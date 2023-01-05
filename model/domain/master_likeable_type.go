package domain

import "online-learning-restful-api/core"

type MasterLikeableType struct {
	core.DomainModel
	LikeableName string `gorm:"varchar(50)"`
}
