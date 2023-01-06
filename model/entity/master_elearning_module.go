package entity

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type MasterElearningModule struct {
	core.EntityModel  `gorm:"embedded"`
	CourseId          uint   `gorm:"not null"`
	ModuleTitle       string `gorm:"type:varchar(100);not null"`
	ModuleOverview    sql.NullString
	ModuleDescription sql.NullString
	IsPublished       bool `gorm:"default:false"`
}

func (MasterElearningModule) TableName() string {
	return "master_elearning_module"
}
