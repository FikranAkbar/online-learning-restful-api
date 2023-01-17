package entity

import (
	"time"
)

type MasterPaymentMethod struct {
	ID        uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"column:name;type:varchar(40);not null"`
}

func (MasterPaymentMethod) TableName() string {
	return "master_payment_method"
}
