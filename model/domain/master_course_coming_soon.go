package domain

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type MasterCourseComingSoon struct {
	core.DomainModel
	Name        string `gorm:"not null"`
	Description sql.NullString
	Cover       sql.NullString
	IsPublished bool `gorm:"default:false"`
}
