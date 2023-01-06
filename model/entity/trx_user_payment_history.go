package entity

import "database/sql"

type TrxUserPaymentHistory struct {
	PaymentId             string               `gorm:"type:varchar(100);primaryKey"`
	UserId                uint                 `gorm:"not null"`
	MasterUser            MasterUser           `gorm:"foreignKey:UserId"`
	CourseId              string               `gorm:"not null"`
	DayId                 uint                 `gorm:"not null"`
	MasterDay             MasterDay            `gorm:"foreignKey:DayId"`
	PaymentChannelId      uint                 `gorm:"not null"`
	MasterPaymentChannel  MasterPaymentChannel `gorm:"foreignKey:PaymentChannelId"`
	PaymentStatusId       uint                 `gorm:"not null"`
	MasterPaymentStatus   MasterPaymentStatus  `gorm:"foreignKey:PaymentStatusId"`
	PromoId               sql.NullInt64
	MasterPromo           MasterPromo `gorm:"foreignKey:PromoId"`
	Price                 uint        `gorm:"not null"`
	TotalPrice            uint        `gorm:"not null"`
	PaymentMethod         uint        `gorm:"not null"`
	PaymentNumber         string
	PaymentInstructionUrl string
	IsExpired             bool `gorm:"default:false"`
}

func (TrxUserPaymentHistory) TableName() string {
	return "trx_user_payment_history"
}
