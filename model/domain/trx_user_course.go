package domain

import "online-learning-restful-api/core"

type TrxUserCourse struct {
	core.DomainModel
	UserId               uint `gorm:"not null"`
	CourseId             uint `gorm:"not null"`
	LastUnlockedModule   uint `gorm:"type:int(10)"`
	TotalWebinarAttended uint `gorm:"type:int(10)"`
	GraduatedAt          string
}
