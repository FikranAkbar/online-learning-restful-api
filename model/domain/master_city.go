package domain

import "online-learning-restful-api/core"

type MasterCity struct {
	core.DomainModel
	CityName   string `gorm:"not null"`
	ProvinceId uint   `gorm:"not null"`
}
