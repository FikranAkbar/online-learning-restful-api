package domain

import "online-learning-restful-api/core"

type MasterCity struct {
	core.DomainModel
	CityName   string `json:"city_name" gorm:"not null" validate:"required"`
	ProvinceId uint   `json:"province_id" gorm:"not null" validate:"required"`
}
