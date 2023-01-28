package course_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/repository/course_repository"
)

type OrderCourseServiceImpl struct {
	course_repository.CourseRepository
	*gorm.DB
}

func NewOrderCourseServiceImpl(courseRepository course_repository.CourseRepository, DB *gorm.DB) *OrderCourseServiceImpl {
	return &OrderCourseServiceImpl{CourseRepository: courseRepository, DB: DB}
}

func (service *OrderCourseServiceImpl) CreateNewCourseOrder(ctx context.Context, courseId uint) course.OrderCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courseOrder, err := service.CourseRepository.CreateNewCourseOrder(ctx, tx, courseId)
	helper.PanicIfError(err)

	return course.OrderCourseResponse{
		UserId:   courseOrder.UserId,
		CourseId: courseOrder.CourseId,
	}
}

func (service *OrderCourseServiceImpl) DeleteCourseOrder(ctx context.Context, courseId uint) course.OrderCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courseOrder, err := service.CourseRepository.DeleteCourseOrder(ctx, tx, courseId)
	helper.PanicIfError(err)

	return course.OrderCourseResponse{
		UserId:   courseOrder.UserId,
		CourseId: courseOrder.CourseId,
	}
}
