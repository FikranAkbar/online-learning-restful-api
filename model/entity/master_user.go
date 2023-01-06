package entity

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type MasterUser struct {
	core.EntityModel `gorm:"embedded"`
	CityId           uint           `gorm:"not null"`
	MasterCity       MasterCity     `gorm:"foreignKey:CityId"`
	ProvinceId       uint           `gorm:"not null"`
	MasterProvince   MasterProvince `gorm:"foreignKey:ProvinceId"`
	Name             string         `gorm:"type:varchar(300);not null"`
	Phone            sql.NullString `gorm:"type:varchar(20)"`
	Gender           string         `gorm:"type:varchar(10);not null"`
	BirthDate        sql.NullTime
	PhotoURL         sql.NullString
}

func (MasterUser) TableName() string {
	return "master_user"
}
