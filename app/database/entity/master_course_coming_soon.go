package entity

import (
	"database/sql"
	"time"
)

type MasterCourseComingSoon struct {
	ID          uint `gorm:"column:id;primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string         `gorm:"column:name;type:varchar(200);not null"`
	Description sql.NullString `gorm:"column:description"`
	Cover       sql.NullString `gorm:"column:cover"`
	IsPublished bool           `gorm:"column:is_published;default:false;not null"`
}

func (MasterCourseComingSoon) TableName() string {
	return "master_course_coming_soon"
}
