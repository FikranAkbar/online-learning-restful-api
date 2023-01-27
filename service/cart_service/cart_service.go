package cart_service

import (
	"context"
	"online-learning-restful-api/model/web/payment"
)

type CartService interface {
	CreateNewCourseOrder(ctx context.Context, courseIds []uint) payment.CourseOrderResponse
}
