package entity

import "online-learning-restful-api/core"

type MasterDay struct {
	core.EntityModel
	DayName string `gorm:"type:varchar(20)"`
}

func (MasterDay) TableName() string {
	return "master_day"
}
