package entity

import (
	"gorm.io/gorm"
)

type TrxUserLike struct {
	gorm.Model     `gorm:"embedded"`
	LikeableId     uint               `gorm:"column:likeable_id;not null"`
	LikeableTypeId uint               `gorm:"column:likeable_type_id;not null"`
	LikeableType   MasterLikeableType `gorm:"foreignKey:LikeableTypeId;joinForeignKey:ID"`
	UserId         uint               `gorm:"column:user_id;not null"`
	User           MasterUser         `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	IsLike         bool               `gorm:"column:is_like;default:false"`
}

func (TrxUserLike) TableName() string {
	return "trx_user_like"
}
