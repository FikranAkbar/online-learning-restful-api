package entity

import (
	"time"
)

type TrxCourseCategory struct {
	CourseId   uint      `gorm:"column:course_id;primaryKey"`
	CategoryId uint      `gorm:"column:category_id;primaryKey"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (TrxCourseCategory) TableName() string {
	return "trx_course_category"
}
