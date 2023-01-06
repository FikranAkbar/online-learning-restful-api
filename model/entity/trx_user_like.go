package entity

import "online-learning-restful-api/core"

type TrxUserLike struct {
	core.EntityModel   `gorm:"embedded"`
	LikeableId         uint               `gorm:"not null"`
	LikeableTypeId     uint               `gorm:"not null"`
	MasterLikeableType MasterLikeableType `gorm:"foreignKey:LikeableTypeId"`
	UserId             uint               `gorm:"not null"`
	MasterUser         MasterUser         `gorm:"foreignKey:UserId"`
	IsLike             bool               `gorm:"default:false"`
}

func (TrxUserLike) TableName() string {
	return "trx_user_like"
}
