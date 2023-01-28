package cart_service

import (
	"context"
	"online-learning-restful-api/model/web/payment"
)

type CartService interface {
	BuyCartItems(ctx context.Context, courseIds []uint) payment.CourseOrderResponse
	HandleMidtransPaymentNotification(ctx context.Context, request payment.MidtransPaymentNotificationRequest)
}
