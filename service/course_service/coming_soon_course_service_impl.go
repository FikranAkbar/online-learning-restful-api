package course_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/repository/course_repository"
)

type ComingSoonCourseServiceImpl struct {
	course_repository.CourseRepository
	*gorm.DB
}

func NewComingSoonCourseServiceImpl(courseRepository course_repository.CourseRepository, DB *gorm.DB) *ComingSoonCourseServiceImpl {
	return &ComingSoonCourseServiceImpl{CourseRepository: courseRepository, DB: DB}
}

func (service *ComingSoonCourseServiceImpl) GetComingSoonCourses(ctx context.Context) []course.ComingSoonCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	comingSoonCourses, err := service.CourseRepository.GetComingSoonCourses(ctx, tx)
	helper.PanicIfError(err)

	var comingSoonCoursesResponse []course.ComingSoonCourseResponse
	for _, comingSoonCourse := range comingSoonCourses {
		comingSoonCoursesResponse = append(comingSoonCoursesResponse, course.ComingSoonCourseResponse{
			Id:         comingSoonCourse.Id,
			CourseName: comingSoonCourse.Name,
			PhotoUrl:   comingSoonCourse.PhotoUrl,
		})
	}

	return comingSoonCoursesResponse
}
