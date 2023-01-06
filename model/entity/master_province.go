package entity

import "online-learning-restful-api/core"

type MasterProvince struct {
	core.EntityModel
	ProvinceName string `gorm:"type:varchar(100)"`
}

func (MasterProvince) TableName() string {
	return "master_province"
}
