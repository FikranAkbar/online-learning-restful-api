package course_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/repository/course_repository"
)

type CourseSummaryServiceImpl struct {
	course_repository.CourseRepository
	*gorm.DB
}

func NewCourseSummaryServiceImpl(courseRepository course_repository.CourseRepository, DB *gorm.DB) *CourseSummaryServiceImpl {
	return &CourseSummaryServiceImpl{CourseRepository: courseRepository, DB: DB}
}

func (service *CourseSummaryServiceImpl) GetCourseSummary(ctx context.Context, courseId uint) course.SummaryResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courseSummary, err := service.CourseRepository.GetCourseSummary(ctx, tx, courseId)
	helper.PanicIfError(err)

	return course.SummaryResponse{
		Id:       courseSummary.ID,
		CourseId: courseSummary.CourseId,
		DocURL:   courseSummary.DocURL,
	}
}
