package entity

import "online-learning-restful-api/core"

type TrxUserCart struct {
	core.EntityModel `gorm:"embedded"`
	UserId           uint `gorm:"column:user_id;not null"`
	CourseId         uint `gorm:"column:course_id;not null"`
}

func (TrxUserCart) TableName() string {
	return "trx_user_cart"
}
