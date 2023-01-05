package domain

import "online-learning-restful-api/core"

type MasterCourse struct {
	core.DomainModel
	ExpertId           uint    `gorm:"not null"`
	StatusId           uint    `gorm:"not null"`
	Name               string  `gorm:"not null"`
	Description        string  `gorm:"not null"`
	PhotoURL           string  `gorm:"not null"`
	AverageRate        float32 `gorm:"default:0.0"`
	Price              uint    `gorm:"not null"`
	TotalRate          uint    `gorm:"default:0"`
	TotalDuration      float32 `gorm:"default:0.0"`
	CurrentParticipant uint    `gorm:"default:0"`
	MaximumParticipant uint    `gorm:"not null"`
	IsPublished        bool    `gorm:"default:false"`
}
