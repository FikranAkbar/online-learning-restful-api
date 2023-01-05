package domain

type TrxUserPaymentHistory struct {
	PaymentId             string `gorm:"type:varchar(100);primaryKey"`
	UserId                uint   `gorm:"not null"`
	CourseId              string `gorm:"not null"`
	DayId                 uint   `gorm:"not null"`
	PaymentChannelId      uint   `gorm:"not null"`
	PaymentStatusId       uint   `gorm:"not null"`
	PromoId               uint   `gorm:"not null"`
	Price                 uint   `gorm:"not null"`
	TotalPrice            uint   `gorm:"not null"`
	PaymentMethod         uint   `gorm:"not null"`
	PaymentNumber         string
	PaymentInstructionUrl string
	IsExpired             bool `gorm:"default:false"`
}
