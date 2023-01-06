package entity

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type MasterCourseComingSoon struct {
	core.EntityModel `gorm:"embedded"`
	Name             string `gorm:"type:varchar(200);not null"`
	Description      sql.NullString
	Cover            sql.NullString
	IsPublished      bool `gorm:"default:false;not null"`
}

func (MasterCourseComingSoon) TableName() string {
	return "master_course_coming_soon"
}
