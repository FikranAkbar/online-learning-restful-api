package entity

import (
	"time"
)

type MasterCategory struct {
	ID           uint `gorm:"column:id;primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CategoryName string         `gorm:"column:category_name;type:varchar(100);not null"`
	Courses      []MasterCourse `gorm:"many2many:trx_course_category;foreignKey:ID;joinForeignKey:CategoryId;references:ID;joinReferences:CourseId"`
}

func (MasterCategory) TableName() string {
	return "master_category"
}
