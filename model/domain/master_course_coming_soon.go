package domain

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type MasterCourseComingSoon struct {
	core.DomainModel
	Name        string `gorm:"type:varchar(200);not null"`
	Description sql.NullString
	Cover       sql.NullString
	IsPublished bool `gorm:"default:false"`
}

func (MasterCourseComingSoon) TableName() string {
	return "master_course_coming_soon"
}
