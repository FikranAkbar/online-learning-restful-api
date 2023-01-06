package entity

import "online-learning-restful-api/core"

type TrxUserCart struct {
	core.EntityModel `gorm:"embedded"`
	UserId           uint `gorm:"not null"`
	CourseId         uint `gorm:"not null"`
}

func (TrxUserCart) TableName() string {
	return "trx_user_cart"
}
