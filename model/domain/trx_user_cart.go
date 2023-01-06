package domain

import "online-learning-restful-api/core"

type TrxUserCart struct {
	core.DomainModel
	UserId   uint `gorm:"not null"`
	CourseId uint `gorm:"not null"`
}

func (TrxUserCart) TableName() string {
	return "trx_user_cart"
}
