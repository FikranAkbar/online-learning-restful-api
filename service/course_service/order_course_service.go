package course_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type OrderCourseService interface {
	CreateNewCourseOrder(ctx context.Context, courseId uint) course.OrderCourseResponse
	DeleteCourseOrder(ctx context.Context, courseId uint) course.OrderCourseResponse
}
