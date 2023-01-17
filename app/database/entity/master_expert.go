package entity

import (
	"database/sql"
	"time"
)

type MasterExpert struct {
	ID                 uint `gorm:"column:id;primarykey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Name               string           `gorm:"column:name;type:varchar(300);not null"`
	Profession         sql.NullString   `gorm:"column:profession;type:varchar(200)"`
	ProfileDescription sql.NullString   `gorm:"column:profession;"`
	Phone              sql.NullString   `gorm:"column:phone;type:varchar(20)"`
	Address            sql.NullString   `gorm:"column:address;"`
	Gender             string           `gorm:"column:gender;type:varchar(20);not null"`
	BirthDate          sql.NullTime     `gorm:"column:birth_date"`
	Photo              sql.NullString   `gorm:"column:photo"`
	AverageRate        float32          `gorm:"column:average_rate;type:float(10);default:0.0"`
	TotalStudent       int              `gorm:"column:total_student;default:0"`
	Categories         []MasterCategory `gorm:"many2many:trx_expert_category;foreignKey:ID;joinForeignKey:ExpertId;references:ID;joinReferences:CategoryId"`
}

func (MasterExpert) TableName() string {
	return "master_expert"
}
