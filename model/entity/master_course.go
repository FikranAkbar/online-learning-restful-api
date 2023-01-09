package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type MasterCourse struct {
	gorm.Model         `gorm:"embedded"`
	ExpertId           uint                    `gorm:"column:expert_id;not null"`
	MasterExpert       MasterExpert            `gorm:"foreignKey:ExpertId;references:ID"`
	StatusId           uint                    `gorm:"column:status_id;not null"`
	MasterCourseStatus MasterCourseStatus      `gorm:"foreignKey:StatusId;references:ID"`
	Name               string                  `gorm:"column:name;type:varchar(100);not null;unique_index"`
	Description        sql.NullString          `gorm:"column:description"`
	PhotoURL           sql.NullString          `gorm:"column:photo_url"`
	AverageRate        float32                 `gorm:"column:average_rate;type:float(10);default:0.0;not null"`
	Price              uint                    `gorm:"column:price;not null"`
	TotalRate          uint                    `gorm:"column:total_rate;default:0;not null"`
	TotalDuration      float32                 `gorm:"column:total_duration;type:float(10);default:0.0;not null"`
	CurrentParticipant uint                    `gorm:"column:current_participant;default:0;not null"`
	MaximumParticipant uint                    `gorm:"column:maximum_participant;default:30;not null"`
	IsPublished        bool                    `gorm:"column:is_published;default:false;not null"`
	Categories         []MasterCategory        `gorm:"many2many:trx_course_category;foreignKey:ID;joinForeignKey:CourseId;references:ID;joinReferences:CategoryId"`
	Videos             []MasterVideo           `gorm:"foreignKey:CourseId;joinForeignKey:ID"`
	Modules            []MasterElearningModule `gorm:"foreignKey:CourseId;joinForeignKey:ID"`
	Webinars           []MasterWebinarSession  `gorm:"foreignKey:CourseId;JoinForeignKey:ID"`
}

func (MasterCourse) TableName() string {
	return "master_course"
}
