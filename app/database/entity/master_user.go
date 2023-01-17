package entity

import (
	"database/sql"
	"time"
)

type MasterUser struct {
	ID               uint `gorm:"column:id;primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CityId           *uint                   `gorm:"column:city_id"`
	MasterCity       MasterCity              `gorm:"foreignKey:CityId"`
	ProvinceId       *uint                   `gorm:"column:province_id"`
	MasterProvince   MasterProvince          `gorm:"foreignKey:ProvinceId"`
	Name             string                  `gorm:"column:name;type:varchar(300);not null"`
	Phone            sql.NullString          `gorm:"column:phone;type:varchar(20)"`
	Gender           string                  `gorm:"column:gender;type:varchar(10);not null"`
	BirthDate        sql.NullTime            `gorm:"column:birth_date"`
	PhotoURL         sql.NullString          `gorm:"column:photo_url"`
	Courses          []MasterCourse          `gorm:"many2many:trx_user_course;foreignKey:ID;joinForeignKey:UserId;references:ID;joinReferences:CourseId"`
	Cart             []TrxUserCart           `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	QuizAnswer       []TrxQuizUserAnswer     `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	PaymentHistories []TrxUserPaymentHistory `gorm:"foreignKey:UserId;joinForeignKey:ID"`
}

func (MasterUser) TableName() string {
	return "master_user"
}
