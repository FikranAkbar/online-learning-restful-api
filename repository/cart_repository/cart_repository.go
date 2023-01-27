package cart_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type CartRepository interface {
	CreateNewCourseOrder(ctx context.Context, db *gorm.DB, courseIds []uint) (domain.PaymentHistory, []domain.Course, domain.User, error)
}
