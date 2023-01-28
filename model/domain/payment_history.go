package domain

type PaymentHistory struct {
	OrderId         string
	UserId          uint
	CourseIds       string
	DayId           uint
	PaymentStatusId uint
	TotalPrice      uint
	PaymentMethodId uint
	PaymentUrl      string
	IsExpired       bool
}
