package domain

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type MasterExpert struct {
	core.DomainModel
	Name               string `gorm:"type:varchar(300)"`
	Profession         string `gorm:"type:varchar(200)"`
	ProfileDescription string
	Phone              string `gorm:"type:varchar(20)"`
	Address            string
	Gender             string `gorm:"type:varchar(20)"`
	BirthDate          sql.NullTime
	Photo              string
	AverageRate        float32 `gorm:"type:float(10)"`
	TotalStudent       uint
}
