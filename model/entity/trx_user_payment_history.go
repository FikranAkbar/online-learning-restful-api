package entity

import "database/sql"

type TrxUserPaymentHistory struct {
	PaymentId             string               `gorm:"column:payment_id;type:varchar(100);primaryKey"`
	UserId                uint                 `gorm:"column:user_id;not null"`
	MasterUser            MasterUser           `gorm:"foreignKey:UserId"`
	CourseId              string               `gorm:"column:course_id;not null"`
	DayId                 uint                 `gorm:"column:day_id;not null"`
	MasterDay             MasterDay            `gorm:"foreignKey:DayId"`
	PaymentChannelId      uint                 `gorm:"column:payment_channel_id;not null"`
	MasterPaymentChannel  MasterPaymentChannel `gorm:"foreignKey:PaymentChannelId"`
	PaymentStatusId       uint                 `gorm:"column:payment_status_id;not null"`
	MasterPaymentStatus   MasterPaymentStatus  `gorm:"foreignKey:PaymentStatusId"`
	PromoId               sql.NullInt64        `gorm:"column:promo_id"`
	MasterPromo           MasterPromo          `gorm:"foreignKey:PromoId"`
	Price                 uint                 `gorm:"column:price;not null"`
	TotalPrice            uint                 `gorm:"column:total_price;not null"`
	PaymentMethod         uint                 `gorm:"column:payment_method;not null"`
	PaymentNumber         string               `gorm:"column:payment_number"`
	PaymentInstructionUrl string               `gorm:"column:payment_instruction_url"`
	IsExpired             bool                 `gorm:"column:is_expired;default:false"`
}

func (TrxUserPaymentHistory) TableName() string {
	return "trx_user_payment_history"
}
