package entity

import (
	"database/sql"
)

type TrxUserCourse struct {
	UserId               uint         `gorm:"column:user_id;primaryKey;autoIncrement:false"`
	User                 MasterUser   `gorm:"foreignKey:UserId"`
	CourseId             uint         `gorm:"column:course_id;primaryKey;autoIncrement:false"`
	Course               MasterCourse `gorm:"foreignKey:CourseId"`
	LastUnlockedModule   uint         `gorm:"column:last_unlocked_module;default=1;not null"`
	TotalWebinarAttended uint         `gorm:"column:total_webinar_attended;default=0;not null"`
	GraduatedAt          sql.NullTime `gorm:"column:graduated_at"`
}

func (TrxUserCourse) TableName() string {
	return "trx_user_course"
}
