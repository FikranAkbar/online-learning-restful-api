package entity

import (
	"database/sql"
	"time"
)

type MasterElearningModule struct {
	ID                uint `gorm:"column:id;primarykey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CourseId          uint                       `gorm:"column:course_id;not null"`
	MasterCourse      MasterCourse               `gorm:"foreignKey:CourseId;references:ID"`
	ModuleTitle       string                     `gorm:"column:module_title;type:varchar(100);not null"`
	ModuleOverview    sql.NullString             `gorm:"column:module_overview;"`
	ModuleDescription sql.NullString             `gorm:"column:module_description"`
	IsPublished       bool                       `gorm:"default:false"`
	Sequence          TrxElearningModuleSequence `gorm:"foreignKey:ModuleId;joinForeignKey:ID"`
	Quiz              MasterQuiz                 `gorm:"foreignKey:ModuleId;joinForeignKey:ID"`
	Video             MasterVideo                `gorm:"foreignKey:ModuleId;joinForeignKey:ID"`
}

func (MasterElearningModule) TableName() string {
	return "master_elearning_module"
}
