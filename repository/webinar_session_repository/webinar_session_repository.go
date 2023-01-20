package webinar_session_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type WebinarSessionRepository interface {
	GetOverviewWebinarSessionsByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.WebinarSession, error)
	GetDetailWebinarSessionByWebinarSessionId(ctx context.Context, db *gorm.DB, courseId uint, webinarSessionId uint) (domain.WebinarSession, error)
}
