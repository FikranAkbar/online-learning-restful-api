package entity

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type TrxUserCourse struct {
	core.EntityModel     `gorm:"embedded"`
	UserId               uint `gorm:"not null"`
	CourseId             uint `gorm:"not null"`
	LastUnlockedModule   uint `gorm:"type:int(10);default=1;not null"`
	TotalWebinarAttended uint `gorm:"type:int(10);default=0;not null"`
	GraduatedAt          sql.NullString
}

func (TrxUserCourse) TableName() string {
	return "trx_user_course"
}
