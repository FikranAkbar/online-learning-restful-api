package user_repository

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"strings"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindUserById(ctx context.Context, db *gorm.DB, userId uint) (domain.User, error) {
	var userEntity entity.MasterUser
	err := db.WithContext(ctx).First(&userEntity, "ID = ?", userId).Error

	if exception.CheckErrorContains(err, exception.NotFound) {
		errorLog := fmt.Sprintf("User data with id %v not found", userId)
		return domain.User{}, exception.GenerateHTTPError(exception.NotFound, errorLog)
	}

	return domain.User{
		Id:   userEntity.ID,
		Name: userEntity.Name,
	}, nil
}

func (repository *UserRepositoryImpl) CreateUserData(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	userEntity := entity.MasterUser{
		ID:     user.Id,
		Name:   user.Name,
		Gender: user.Gender,
		BirthDate: sql.NullTime{
			Time:  user.BirthDate,
			Valid: true,
		},
	}
	err := db.WithContext(ctx).Create(&userEntity).Error
	helper.PanicIfError(err)

	return domain.User{
		Id:        userEntity.ID,
		Name:      userEntity.Name,
		BirthDate: userEntity.BirthDate.Time,
		Gender:    userEntity.Gender,
	}, nil
}

func (repository *UserRepositoryImpl) GetUserCourses(ctx context.Context, db *gorm.DB, courseStatus string) ([]domain.Course, error) {
	if courseStatus == "" {
		return nil, exception.GenerateHTTPError(exception.BadRequest, "Type query can't be empty")
	}

	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	var userEntity entity.MasterUser
	err := db.WithContext(ctx).
		Where("id = ?", userTokenInfo.UserId).
		Preload("Courses.Expert").
		Preload("Courses.CourseStatus").
		First(&userEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("User with id %v not found", userTokenInfo.UserId)
		return nil, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return nil, err
	}

	courseEntities := userEntity.Courses
	var courses []domain.Course
	for _, courseEntity := range courseEntities {
		if !courseEntity.IsPublished {
			continue
		}

		if strings.ToLower(courseEntity.CourseStatus.Name) !=
			strings.ToLower(courseStatus) {
			continue
		}

		courses = append(courses, domain.Course{
			Id:          courseEntity.ID,
			ExpertName:  courseEntity.Expert.Name,
			Name:        courseEntity.Name,
			PhotoUrl:    courseEntity.PhotoURL.String,
			AverageRate: courseEntity.AverageRate,
		})
	}

	return courses, nil
}

func (repository *UserRepositoryImpl) GetProvinces(ctx context.Context, db *gorm.DB) ([]domain.Province, error) {
	var provinceEntities []entity.MasterProvince
	err := db.WithContext(ctx).Find(&provinceEntities).Error
	if err != nil {
		return nil, err
	}

	var provinces []domain.Province
	for _, provinceEntity := range provinceEntities {
		provinces = append(provinces, domain.Province{
			Id:           provinceEntity.ID,
			ProvinceName: provinceEntity.ProvinceName,
		})
	}

	return provinces, nil
}

func (repository *UserRepositoryImpl) GetCitiesByProvinceId(ctx context.Context, db *gorm.DB, provinceId uint) ([]domain.City, error) {
	var cityEntities []entity.MasterCity
	err := db.WithContext(ctx).
		Where("province_id = ?", provinceId).
		Find(&cityEntities).Error
	if err != nil {
		return nil, err
	}

	var cities []domain.City
	for _, cityEntity := range cityEntities {
		cities = append(cities, domain.City{
			Id:         cityEntity.ID,
			CityName:   cityEntity.CityName,
			ProvinceId: cityEntity.ProvinceId,
		})
	}

	return cities, nil
}
