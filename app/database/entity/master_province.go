package entity

import "gorm.io/gorm"

type MasterProvince struct {
	gorm.Model   `gorm:"embedded"`
	ProvinceName string `gorm:"column:province_name;type:varchar(100);not null"`
}

func (MasterProvince) TableName() string {
	return "master_province"
}
