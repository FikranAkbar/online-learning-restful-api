package industry_insight_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/industry_insight"
	"online-learning-restful-api/repository/industry_insight_repository"
)

type IndustryInsightServiceImpl struct {
	industry_insight_repository.IndustryInsightRepository
	*gorm.DB
}

func NewIndustryInsightServiceImpl(industryInsightRepository industry_insight_repository.IndustryInsightRepository, DB *gorm.DB) *IndustryInsightServiceImpl {
	return &IndustryInsightServiceImpl{IndustryInsightRepository: industryInsightRepository, DB: DB}
}

func (service *IndustryInsightServiceImpl) GetIndustryInsights(ctx context.Context) []industry_insight.IndustryInsightResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	industryInsights, err := service.IndustryInsightRepository.GetIndustryInsights(ctx, tx)
	helper.PanicIfError(err)

	var industryInsightsResponse []industry_insight.IndustryInsightResponse
	for _, industryInsight := range industryInsights {
		industryInsightsResponse = append(industryInsightsResponse, industry_insight.IndustryInsightResponse{
			Id:           industryInsight.Id,
			TitleInsight: industryInsight.TitleInsight,
			CoverInsight: industryInsight.CoverInsight,
			BodyContent:  industryInsight.BodyContent,
			IdUserAuthor: industryInsight.UserAuthorId,
		})
	}

	return industryInsightsResponse
}

func (service *IndustryInsightServiceImpl) GetIndustryInsightById(ctx context.Context, industryInsightId uint) industry_insight.IndustryInsightResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	industryInsight, insightAttachments, err := service.IndustryInsightRepository.GetIndustryInsightById(ctx, tx, industryInsightId)
	helper.PanicIfError(err)

	var insightAttachmentsResponse []industry_insight.InsightAttachmentResponse
	for _, insightAttachment := range insightAttachments {
		insightAttachmentsResponse = append(insightAttachmentsResponse, industry_insight.InsightAttachmentResponse{
			Id:             insightAttachment.ID,
			AttachmentName: insightAttachment.AttachmentName,
			DocURL:         insightAttachment.DocURL,
		})
	}

	return industry_insight.IndustryInsightResponse{
		Id:                 industryInsight.Id,
		TitleInsight:       industryInsight.TitleInsight,
		CoverInsight:       industryInsight.CoverInsight,
		BodyContent:        industryInsight.BodyContent,
		IdUserAuthor:       industryInsight.UserAuthorId,
		InsightAttachments: insightAttachmentsResponse,
	}
}
