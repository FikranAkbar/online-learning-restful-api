package entity

import (
	"time"
)

type TrxCourseReview struct {
	ID           uint `gorm:"column:id;primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CourseId     uint         `gorm:"column:course_id;not null"`
	MasterCourse MasterCourse `gorm:"foreignKey:CourseId;references:ID"`
	UserId       uint         `gorm:"column:user_id;not null"`
	MasterUser   MasterUser   `gorm:"foreignKey:UserId;references:ID"`
	ReviewDesc   string       `gorm:"column:review_desc;not null"`
	ReviewRate   float32      `gorm:"column:review_rate;not null"`
}

func (TrxCourseReview) TableName() string {
	return "trx_course_review"
}
