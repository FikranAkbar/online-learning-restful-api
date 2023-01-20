package course_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type CourseReviewService interface {
	GetCourseReviewsByCourseId(ctx context.Context, courseId uint) []course.ReviewCourseResponse
}
