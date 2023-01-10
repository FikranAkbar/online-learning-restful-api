package entity

import (
	"database/sql"
)

type TrxUserCourse struct {
	UserId               uint           `gorm:"column:user_id;primaryKey"`
	CourseId             uint           `gorm:"column:course_id;primaryKey"`
	LastUnlockedModule   uint           `gorm:"column:last_unlocked_module;type:int(10);default=1;not null"`
	TotalWebinarAttended uint           `gorm:"column:total_webinar_attended;type:int(10);default=0;not null"`
	GraduatedAt          sql.NullString `gorm:"column:graduated_at"`
}

func (TrxUserCourse) TableName() string {
	return "trx_user_course"
}