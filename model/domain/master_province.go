package domain

import "online-learning-restful-api/core"

type MasterProvince struct {
	core.DomainModel
	ProvinceName string `gorm:"type:varchar(100)"`
}
