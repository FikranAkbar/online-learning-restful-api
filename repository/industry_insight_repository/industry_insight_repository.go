package industry_insight_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type IndustryInsightRepository interface {
	GetIndustryInsights(ctx context.Context, db *gorm.DB) ([]domain.IndustryInsight, error)
}
