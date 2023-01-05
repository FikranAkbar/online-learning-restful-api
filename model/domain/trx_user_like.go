package domain

import "online-learning-restful-api/core"

type TrxUserLike struct {
	core.DomainModel
	LikeableId     uint `gorm:"not null"`
	LikeableTypeId uint `gorm:"not null"`
	UserId         uint `gorm:"not null"`
	IsLike         bool `gorm:"default:false"`
}
