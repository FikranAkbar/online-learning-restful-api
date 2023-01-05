package domain

import "online-learning-restful-api/core"

type MasterUser struct {
	core.DomainModel
	CityId     uint   `gorm:"not null"`
	ProvinceId uint   `gorm:"not null"`
	Name       string `gorm:"type:varchar(300)"`
	Phone      string `gorm:"type:varchar(20)"`
	Gender     string `gorm:"type:varchar(10)"`
	BirthDate  string
	PhotoURL   string
}
