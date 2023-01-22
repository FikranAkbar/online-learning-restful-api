package industry_insight_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/model/domain"
)

type IndustryInsightRepositoryImpl struct {
}

func NewIndustryInsightRepositoryImpl() *IndustryInsightRepositoryImpl {
	return &IndustryInsightRepositoryImpl{}
}

func (repository *IndustryInsightRepositoryImpl) GetIndustryInsights(ctx context.Context, db *gorm.DB) ([]domain.IndustryInsight, error) {
	var industryInsightEntities []entity.MasterIndustryInsight
	err := db.WithContext(ctx).Find(&industryInsightEntities).Error
	if err != nil {
		return []domain.IndustryInsight{}, err
	}

	var industryInsights []domain.IndustryInsight
	for _, industryInsightEntity := range industryInsightEntities {
		if !industryInsightEntity.IsPublished {
			continue
		}
		industryInsights = append(industryInsights, domain.IndustryInsight{
			Id:           industryInsightEntity.ID,
			UserAuthorId: industryInsightEntity.UserAuthorId,
			TitleInsight: industryInsightEntity.TitleInsight,
			CoverInsight: industryInsightEntity.CoverInsight.String,
			BodyContent:  industryInsightEntity.BodyContent.String,
			IsPublished:  industryInsightEntity.IsPublished,
		})
	}

	return industryInsights, nil
}

func (repository *IndustryInsightRepositoryImpl) GetIndustryInsightById(ctx context.Context, db *gorm.DB, industryInsightId uint) (domain.IndustryInsight, []domain.InsightAttachment, error) {
	var industryInsightEntity entity.MasterIndustryInsight
	err := db.WithContext(ctx).
		Where("id = ?", industryInsightId).
		Preload("Attachments").
		First(&industryInsightEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) || !industryInsightEntity.IsPublished {
		logError := fmt.Sprintf("Industry insights with id %v not found", industryInsightId)
		return domain.IndustryInsight{}, []domain.InsightAttachment{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.IndustryInsight{}, []domain.InsightAttachment{}, err
	}

	insightAttachmentEntities := industryInsightEntity.Attachments

	industryInsight := domain.IndustryInsight{
		Id:           industryInsightEntity.ID,
		UserAuthorId: industryInsightEntity.UserAuthorId,
		TitleInsight: industryInsightEntity.TitleInsight,
		CoverInsight: industryInsightEntity.CoverInsight.String,
		BodyContent:  industryInsightEntity.BodyContent.String,
		IsPublished:  industryInsightEntity.IsPublished,
	}

	var insightAttachments []domain.InsightAttachment
	for _, insightAttachment := range insightAttachmentEntities {
		insightAttachments = append(insightAttachments, domain.InsightAttachment{
			ID:             insightAttachment.ID,
			AttachmentName: insightAttachment.AttachmentName,
			DocURL:         insightAttachment.DocURL,
		})
	}

	return industryInsight, insightAttachments, nil
}
