package domain

import "online-learning-restful-api/core"

type MasterCategory struct {
	core.DomainModel
	CategoryName string `json:"category_name" gorm:"not null" validate:"required"`
}
