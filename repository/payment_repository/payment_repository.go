package payment_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type PaymentRepository interface {
	CreateNewCourseOrder(ctx context.Context, db *gorm.DB, courseIds []uint) (domain.User, []domain.Course, error)
}
