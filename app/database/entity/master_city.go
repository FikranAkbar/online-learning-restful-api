package entity

import (
	"time"
)

type MasterCity struct {
	ID         uint `gorm:"column:id;primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CityName   string         `gorm:"column:city_name;type:varchar(100);not null"`
	ProvinceId uint           `gorm:"column:province_id"`
	Province   MasterProvince `gorm:"foreignKey:ProvinceId;references:ID"`
}

func (MasterCity) TableName() string {
	return "master_city"
}
