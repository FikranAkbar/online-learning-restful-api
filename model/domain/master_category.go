package domain

import "online-learning-restful-api/core"

type MasterCategory struct {
	core.DomainModel
	CategoryName string `gorm:"type:varchar(100);not null"`
}
