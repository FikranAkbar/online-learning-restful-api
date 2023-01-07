package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type MasterCourseComingSoon struct {
	gorm.Model  `gorm:"embedded"`
	Name        string         `gorm:"column:name;type:varchar(200);not null"`
	Description sql.NullString `gorm:"column:description"`
	Cover       sql.NullString `gorm:"column:cover"`
	IsPublished bool           `gorm:"column:is_published;default:false;not null"`
}

func (MasterCourseComingSoon) TableName() string {
	return "master_course_coming_soon"
}
