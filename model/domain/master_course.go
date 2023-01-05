package domain

import "online-learning-restful-api/core"

type MasterCourse struct {
	core.DomainModel
	ExpertId           uint    `json:"expert_id" gorm:"not null" validate:"required"`
	StatusId           uint    `json:"status_id" gorm:"not null" validate:"required"`
	Name               string  `json:"name" gorm:"not null" validate:"required"`
	Description        string  `json:"description" gorm:"not null" validate:"required"`
	PhotoURL           string  `json:"photo_url" gorm:"not null" validate:"required"`
	AverageRate        float32 `json:"average_rate" gorm:"default:0.0"`
	Price              uint    `json:"price" gorm:"not null" validate:"required"`
	TotalRate          uint    `json:"total_rate" gorm:"default:0.0" validate:"required"`
	TotalDuration      float32 `json:"total_duration" gorm:"default:0.0" validate:"required"`
	CurrentParticipant uint    `json:"current_participant" gorm:"default:0.0" validate:"required"`
	MaximumParticipant uint    `json:"maximum_participant" gorm:"not null" validate:"required"`
	IsPublished        bool    `json:"is_published" gorm:"default:false" validate:"required"`
}
