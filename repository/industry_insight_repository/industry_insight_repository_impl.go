package industry_insight_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
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
