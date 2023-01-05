package domain

import "online-learning-restful-api/core"

type name struct {
	core.DomainModel
	DayName string `gorm:"type:varchar(20)"`
}
