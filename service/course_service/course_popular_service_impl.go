package course_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/repository/course_repository"
)

type CoursePopularServiceImpl struct {
	course_repository.CourseRepository
	*gorm.DB
}

func NewCoursePopularServiceImpl(repository course_repository.CourseRepository, db *gorm.DB) *CoursePopularServiceImpl {
	return &CoursePopularServiceImpl{
		CourseRepository: repository,
		DB:               db,
	}
}

func (service *CoursePopularServiceImpl) GetPopularCourses(ctx context.Context) []course.ShortCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courses, err := service.CourseRepository.GetPopularCourses(ctx, tx)
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
