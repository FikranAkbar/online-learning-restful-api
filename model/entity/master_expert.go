package entity

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type MasterExpert struct {
	core.EntityModel   `gorm:"embedded"`
	Name               string         `gorm:"type:varchar(300);not null"`
	Profession         sql.NullString `gorm:"type:varchar(200)"`
	ProfileDescription sql.NullString
	Phone              sql.NullString `gorm:"type:varchar(20)"`
	Address            sql.NullString
	Gender             string `gorm:"type:varchar(20);not null"`
	BirthDate          sql.NullTime
	Photo              sql.NullString
	AverageRate        float32 `gorm:"type:float(10);default:0.0"`
	TotalStudent       int     `gorm:"default:0"`
}

func (MasterExpert) TableName() string {
	return "master_expert"
}
