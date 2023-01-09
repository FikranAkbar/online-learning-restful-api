package entity

import (
	"gorm.io/gorm"
	"time"
)

type TrxCourseCategory struct {
	CourseId   uint           `gorm:"column:course_id;primaryKey"`
	CategoryId uint           `gorm:"column:category_id;primaryKey"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (TrxCourseCategory) TableName() string {
	return "trx_course_category"
}
