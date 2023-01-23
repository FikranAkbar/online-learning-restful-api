package user_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type UserService interface {
	GetUserCourses(ctx context.Context, courseStatus string) []course.ShortCourseResponse
}
