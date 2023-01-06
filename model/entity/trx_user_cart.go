package entity

import "online-learning-restful-api/core"

type TrxUserCart struct {
	core.EntityModel `gorm:"embedded"`
	UserId           uint         `gorm:"not null"`
	MasterUser       MasterUser   `gorm:"foreignKey:UserId"`
	CourseId         uint         `gorm:"not null"`
	MasterCourse     MasterCourse `gorm:"foreignKey:CourseId"`
}

func (TrxUserCart) TableName() string {
	return "trx_user_cart"
}
