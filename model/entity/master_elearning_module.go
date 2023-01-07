package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type MasterElearningModule struct {
	gorm.Model        `gorm:"embedded"`
	CourseId          uint           `gorm:"column:course_id;not null"`
	MasterCourse      MasterCourse   `gorm:"foreignKey:CourseId;references:ID"`
	ModuleTitle       string         `gorm:"column:module_title;type:varchar(100);not null"`
	ModuleOverview    sql.NullString `gorm:"column:module_overview;"`
	ModuleDescription sql.NullString `gorm:"column:module_description"`
	IsPublished       bool           `gorm:"default:false"`
}

func (MasterElearningModule) TableName() string {
	return "master_elearning_module"
}
