package cart_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type CartRepository interface {
	BuyCartItems(ctx context.Context, db *gorm.DB, courseIds []uint) (domain.PaymentHistory, []domain.Course, domain.User, error)
	SavePaymentRedirectionURL(ctx context.Context, db *gorm.DB, paymentHistory domain.PaymentHistory) (domain.PaymentHistory, error)
	AddNewCourseToUser(ctx context.Context, db *gorm.DB, paymentNotification domain.MidtransPaymentNotification) error
	ChangeUserPaymentHistoryToCancelled(ctx context.Context, db *gorm.DB, paymentNotification domain.MidtransPaymentNotification) error
}
