package domain

import "online-learning-restful-api/core"

type MasterCity struct {
	core.DomainModel
	CityName   string `gorm:"type:varchar(100);not null"`
	ProvinceId uint   `gorm:"not null"`
}

func (MasterCity) TableName() string {
	return "master_city"
}
