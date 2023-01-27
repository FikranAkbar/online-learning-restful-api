package payment_service

import (
	"context"
	"online-learning-restful-api/model/web/payment"
)

type PaymentService interface {
	CreateNewCourseOrder(ctx context.Context, courseIds []int) payment.CourseOrderResponse
}
