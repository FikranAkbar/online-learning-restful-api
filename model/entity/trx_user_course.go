package entity

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type TrxUserCourse struct {
	core.EntityModel     `gorm:"embedded"`
	UserId               uint               `gorm:"not null"`
	MasterUser           MasterUser         `gorm:"foreignKey:UserId"`
	CourseId             uint               `gorm:"not null"`
	MasterCourse         MasterCourseStatus `gorm:"foreignKey:CourseId"`
	LastUnlockedModule   uint               `gorm:"type:int(10);default=1;not null"`
	TotalWebinarAttended uint               `gorm:"type:int(10);default=0;not null"`
	GraduatedAt          sql.NullString
}

func (TrxUserCourse) TableName() string {
	return "trx_user_course"
}
