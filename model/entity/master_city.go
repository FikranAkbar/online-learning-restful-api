package entity

import (
	"gorm.io/gorm"
)

type MasterCity struct {
	gorm.Model     `gorm:"embedded"`
	CityName       string         `gorm:"column:city_name;type:varchar(100);not null"`
	ProvinceId     uint           `gorm:"column:province_id"`
	MasterProvince MasterProvince `gorm:"foreignKey:ProvinceId;references:ID"`
}

func (MasterCity) TableName() string {
	return "master_city"
}
