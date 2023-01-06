package entity

import "online-learning-restful-api/core"

type MasterProvince struct {
	core.EntityModel `gorm:"embedded"`
	ProvinceName     string `gorm:"type:varchar(100);not null"`
}

func (MasterProvince) TableName() string {
	return "master_province"
}
