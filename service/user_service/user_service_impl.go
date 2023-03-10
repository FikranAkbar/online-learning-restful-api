package user_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"online-learning-restful-api/model/web/city"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/model/web/province"
	"online-learning-restful-api/model/web/user"
	"online-learning-restful-api/repository/user_repository"
)

type UserServiceImpl struct {
	user_repository.UserRepository
	*gorm.DB
}

func NewUserServiceImpl(userRepository user_repository.UserRepository, DB *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB}
}

func (service *UserServiceImpl) GetUserCourses(ctx context.Context, courseStatus string) []course.ShortCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courses, err := service.UserRepository.GetUserCourses(ctx, tx, courseStatus)
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

func (service *UserServiceImpl) GetProvinces(ctx context.Context) []province.ShortProvinceResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	provinces, err := service.UserRepository.GetProvinces(ctx, tx)
	helper.PanicIfError(err)

	var provincesResponse []province.ShortProvinceResponse
	for _, provinceData := range provinces {
		provincesResponse = append(provincesResponse, province.ShortProvinceResponse{
			Id:           provinceData.Id,
			ProvinceName: provinceData.ProvinceName,
		})
	}

	return provincesResponse
}

func (service *UserServiceImpl) GetCitiesByProvinceId(ctx context.Context, provinceId uint) []city.ShortCityResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	cities, err := service.UserRepository.GetCitiesByProvinceId(ctx, tx, provinceId)
	helper.PanicIfError(err)

	var citiesResponse []city.ShortCityResponse
	for _, cityData := range cities {
		citiesResponse = append(citiesResponse, city.ShortCityResponse{
			Id:         cityData.Id,
			CityName:   cityData.CityName,
			ProvinceId: cityData.ProvinceId,
		})
	}

	return citiesResponse
}

func (service *UserServiceImpl) EditUserProfile(ctx context.Context, request user.EditProfileRequest) user.DetailUserResponse {
	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	userDomain := domain.User{
		Id:           userTokenInfo.UserId,
		Name:         request.Name,
		Gender:       request.Gender,
		Phone:        request.Phone,
		PhotoURL:     request.Photo,
		ProvinceName: request.ProvinceName,
	}

	userData, err := service.UserRepository.EditUserProfile(ctx, tx, userDomain)
	helper.PanicIfError(err)

	return user.DetailUserResponse{
		Id:         userData.Id,
		Name:       userData.Name,
		BirthDate:  userData.BirthDate,
		Gender:     userData.Gender,
		Phone:      userData.Phone,
		PhotoURL:   userData.PhotoURL,
		ProvinceId: userData.ProvinceId,
	}
}
