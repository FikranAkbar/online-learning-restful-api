package entity

import "online-learning-restful-api/core"

type MasterCourse struct {
	core.EntityModel   `gorm:"embedded"`
	ExpertId           uint               `gorm:"not null"`
	MasterExpert       MasterExpert       `gorm:"foreignKey:ExpertId"`
	StatusId           uint               `gorm:"not null"`
	MasterCourseStatus MasterCourseStatus `gorm:"foreignKey:StatusId"`
	Name               string             `gorm:"type:varchar(100);not null;unique_index"`
	Description        string             `gorm:"not null"`
	PhotoURL           string             `gorm:"not null"`
	AverageRate        float32            `gorm:"type:float(10);default:0.0;not null"`
	Price              uint               `gorm:"not null"`
	TotalRate          uint               `gorm:"default:0"`
	TotalDuration      float32            `gorm:"type:float(10);default:0.0;not null"`
	CurrentParticipant uint               `gorm:"default:0;not null"`
	MaximumParticipant uint               `gorm:"not null"`
	IsPublished        bool               `gorm:"default:false;not null"`
}

func (MasterCourse) TableName() string {
	return "master_course"
}
