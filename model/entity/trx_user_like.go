package entity

import "online-learning-restful-api/core"

type TrxUserLike struct {
	core.EntityModel `gorm:"embedded"`
	LikeableId       uint `gorm:"column:likeable_id;not null"`
	LikeableTypeId   uint `gorm:"column:likeable_type_id;not null"`
	UserId           uint `gorm:"column:user_id;not null"`
	IsLike           bool `gorm:"column:is_like;default:false"`
}

func (TrxUserLike) TableName() string {
	return "trx_user_like"
}
