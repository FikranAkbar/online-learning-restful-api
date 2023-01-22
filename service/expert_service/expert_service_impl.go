package expert_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/expert"
	"online-learning-restful-api/repository/expert_repository"
)

type ExpertServiceImpl struct {
	expert_repository.ExpertRepository
	*gorm.DB
}

func NewExpertServiceImpl(expertRepository expert_repository.ExpertRepository, DB *gorm.DB) *ExpertServiceImpl {
	return &ExpertServiceImpl{ExpertRepository: expertRepository, DB: DB}
}

func (service *ExpertServiceImpl) GetExpertDetailById(ctx context.Context, expertId uint) expert.DetailExpertResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	expertDetail, err := service.ExpertRepository.GetExpertDetailById(ctx, tx, expertId)
	helper.PanicIfError(err)

	return expert.DetailExpertResponse{
		ExpertId:     expertDetail.Id,
		Name:         expertDetail.Name,
		Photo:        expertDetail.Photo,
		AverageRate:  expertDetail.AverageRate,
		TotalStudent: expertDetail.TotalStudent,
		TotalCourse:  expertDetail.TotalCourse,
	}
}
