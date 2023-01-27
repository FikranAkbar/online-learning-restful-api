package entity

type TrxUserPaymentHistory struct {
	OrderId         string              `gorm:"column:order_id;type:varchar(100);primaryKey"`
	UserId          uint                `gorm:"column:user_id;not null"`
	CourseIds       string              `gorm:"column:course_ids;not null"`
	DayId           uint                `gorm:"column:day_id;not null"`
	Day             MasterDay           `gorm:"foreignKey:DayId;joinForeignKey:ID"`
	PaymentStatusId uint                `gorm:"column:payment_status_id;not null"`
	PaymentStatus   MasterPaymentStatus `gorm:"foreignKey:PaymentStatusId;joinForeignKey:ID"`
	PromoId         *uint               `gorm:"column:promo_id;default:null"`
	Promo           MasterPromo         `gorm:"foreignKey:PromoId;joinForeignKey:ID"`
	Price           uint                `gorm:"column:price;not null"`
	TotalPrice      uint                `gorm:"column:total_price;not null"`
	PaymentMethodId uint                `gorm:"column:payment_method;not null"`
	PaymentMethod   MasterPaymentMethod `gorm:"foreignKey:PaymentMethodId;joinForeignKey:ID"`
	PaymentUrl      string              `gorm:"column:payment_url"`
	IsExpired       bool                `gorm:"column:is_expired;default:false"`
}

func (TrxUserPaymentHistory) TableName() string {
	return "trx_user_payment_history"
}
