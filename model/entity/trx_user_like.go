package entity

import "online-learning-restful-api/core"

type TrxUserLike struct {
	core.EntityModel `gorm:"embedded"`
	LikeableId       uint `gorm:"not null"`
	LikeableTypeId   uint `gorm:"not null"`
	UserId           uint `gorm:"not null"`
	IsLike           bool `gorm:"default:false"`
}

func (TrxUserLike) TableName() string {
	return "trx_user_like"
}
