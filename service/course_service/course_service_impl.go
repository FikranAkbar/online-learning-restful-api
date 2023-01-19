package course_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/repository/course_repository"
)

type CourseServiceImpl struct {
	course_repository.CourseRepository
	*gorm.DB
}

func NewCourseServiceImpl(courseRepository course_repository.CourseRepository, DB *gorm.DB) *CourseServiceImpl {
	return &CourseServiceImpl{CourseRepository: courseRepository, DB: DB}
}

func (service *CourseServiceImpl) GetCoursesByKeyword(ctx context.Context, keyword string) []course.ShortCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courses, err := service.CourseRepository.GetCoursesByKeyword(ctx, tx, keyword)
	helper.PanicIfError(err)

	var coursesResponse []course.ShortCourseResponse
	for _, courseData := range courses {
		coursesResponse = append(coursesResponse, course.ShortCourseResponse{
			Id:          courseData.Id,
			Name:        courseData.Name,
			PhotoUrl:    courseData.PhotoUrl,
			AverageRate: courseData.AverageRate,
			ExpertName:  courseData.ExpertName,
		})
	}

	return coursesResponse
}
