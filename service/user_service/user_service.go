package user_service

import (
	"context"
	"online-learning-restful-api/model/web/city"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/model/web/province"
	"online-learning-restful-api/model/web/user"
)

type UserService interface {
	GetUserCourses(ctx context.Context, courseStatus string) []course.ShortCourseResponse
	GetProvinces(ctx context.Context) []province.ShortProvinceResponse
	GetCitiesByProvinceId(ctx context.Context, provinceId uint) []city.ShortCityResponse
	EditUserProfile(ctx context.Context, request user.EditProfileRequest) user.DetailUserResponse
}
