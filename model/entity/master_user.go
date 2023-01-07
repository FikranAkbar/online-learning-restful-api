package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type MasterUser struct {
	gorm.Model     `gorm:"embedded"`
	CityId         uint           `gorm:"column:city_id;not null"`
	MasterCity     MasterCity     `gorm:"foreignKey:CityId"`
	ProvinceId     uint           `gorm:"column:province_id;not null"`
	MasterProvince MasterProvince `gorm:"foreignKey:ProvinceId"`
	Name           string         `gorm:"column:name;type:varchar(300);not null"`
	Phone          sql.NullString `gorm:"column:phone;type:varchar(20)"`
	Gender         string         `gorm:"column:gender;type:varchar(10);not null"`
	BirthDate      sql.NullTime   `gorm:"column:birth_date"`
	PhotoURL       sql.NullString `gorm:"column:photo_url"`
}

func (MasterUser) TableName() string {
	return "master_user"
}
