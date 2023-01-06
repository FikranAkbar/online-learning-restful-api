package entity

import "online-learning-restful-api/core"

type TrxUserCourse struct {
	core.EntityModel
	UserId               uint `gorm:"not null"`
	CourseId             uint `gorm:"not null"`
	LastUnlockedModule   uint `gorm:"type:int(10)"`
	TotalWebinarAttended uint `gorm:"type:int(10)"`
	GraduatedAt          string
}

func (TrxUserCourse) TableName() string {
	return "trx_user_course"
}
