package industry_insight_service

import (
	"context"
	"online-learning-restful-api/model/web/industry_insight"
)

type IndustryInsightService interface {
	GetIndustryInsights(ctx context.Context) []industry_insight.IndustryInsightResponse
	GetIndustryInsightById(ctx context.Context, industryInsightId uint) industry_insight.IndustryInsightResponse
}
