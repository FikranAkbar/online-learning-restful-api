package expert_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
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

func (service *ExpertServiceImpl) GetExpertCoursesById(ctx context.Context, expertId uint) []course.ShortCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courses, err := service.ExpertRepository.GetExpertCoursesById(ctx, tx, expertId)
	helper.PanicIfError(err)

	var coursesResponse []course.ShortCourseResponse
	for _, data := range courses {
		coursesResponse = append(coursesResponse, course.ShortCourseResponse{
			Id:          data.Id,
			Name:        data.Name,
			PhotoUrl:    data.PhotoUrl,
			AverageRate: data.AverageRate,
			ExpertName:  data.ExpertName,
		})
	}

	return coursesResponse
}
