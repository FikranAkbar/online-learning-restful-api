package course_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"strings"
)

type CourseRepositoryImpl struct {
}

func NewCourseRepositoryImpl() *CourseRepositoryImpl {
	return &CourseRepositoryImpl{}
}

func (repository *CourseRepositoryImpl) GetPopularCourses(ctx context.Context, db *gorm.DB) ([]domain.Course, error) {
	var courseEntities []entity.MasterCourse
	err := db.WithContext(ctx).Limit(5).Preload("Expert").Find(&courseEntities).Error
	if err != nil {
		return []domain.Course{}, err
	}

	var courses []domain.Course
	for _, courseEntity := range courseEntities {
		courses = append(courses, domain.Course{
			Id:          courseEntity.ID,
			Name:        courseEntity.Name,
			PhotoUrl:    courseEntity.PhotoURL.String,
			AverageRate: courseEntity.AverageRate,
			ExpertName:  courseEntity.Expert.Name,
		})
	}

	return courses, nil
}

func (repository *CourseRepositoryImpl) GetCoursesByKeyword(ctx context.Context, db *gorm.DB, keyword string) ([]domain.Course, error) {
	var courseEntities []entity.MasterCourse
	newKeyword := "%" + keyword + "%"
	err := db.WithContext(ctx).Preload("Expert").Where("Name LIKE ?", newKeyword).Find(&courseEntities).Error
	if err != nil {
		return []domain.Course{}, err
	}

	var courses []domain.Course
	for _, courseEntity := range courseEntities {
		courses = append(courses, domain.Course{
			Id:          courseEntity.ID,
			Name:        courseEntity.Name,
			PhotoUrl:    courseEntity.PhotoURL.String,
			AverageRate: courseEntity.AverageRate,
			ExpertName:  courseEntity.Expert.Name,
		})
	}

	return courses, nil
}

func (repository *CourseRepositoryImpl) GetCourseDetailByCourseId(ctx context.Context, db *gorm.DB, courseId uint, userId *uint) (domain.Course, domain.Expert, error) {
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Preload("Expert").
		Preload("CourseStatus").
		Preload("Modules").
		Preload("Webinars").
		First(&courseEntity, courseId).Error
	if err != nil {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return domain.Course{}, domain.Expert{}, exception.GenerateHTTPError(exception.NotFound, logError)
	}

	expertEntity := courseEntity.Expert
	err = db.WithContext(ctx).Preload("Reviews").Error
	helper.PanicIfError(err)

	alreadyOwned := false
	if userId != nil {
		var userEntity entity.MasterUser
		err = db.WithContext(ctx).First(&userEntity, userId).Error
		if exception.CheckErrorContains(err, strings.ToLower(exception.NotFound)) {
			alreadyOwned = false
		} else {
			helper.PanicIfError(err)
		}
	}

	course := domain.Course{
		Id:                 courseEntity.ID,
		ExpertId:           courseEntity.ExpertId,
		ExpertName:         courseEntity.Expert.Name,
		StatusCourse:       courseEntity.CourseStatus.Name,
		Name:               courseEntity.Name,
		Description:        courseEntity.Description.String,
		PhotoUrl:           courseEntity.PhotoURL.String,
		AverageRate:        courseEntity.AverageRate,
		Price:              courseEntity.Price,
		TotalRate:          courseEntity.TotalRate,
		TotalDuration:      courseEntity.TotalDuration,
		CurrentParticipant: courseEntity.CurrentParticipant,
		MaximumParticipant: courseEntity.MaximumParticipant,
		AlreadyOwned:       alreadyOwned,
		ModulesCount:       len(courseEntity.Modules),
		WebinarsCount:      len(courseEntity.Webinars),
	}

	expert := domain.Expert{
		Id:           expertEntity.ID,
		Name:         expertEntity.Name,
		Profession:   expertEntity.Profession.String,
		Phone:        expertEntity.Phone.String,
		Address:      expertEntity.Address.String,
		Gender:       expertEntity.Gender,
		Photo:        expertEntity.Photo.String,
		BirthDate:    expertEntity.BirthDate.Time,
		AverageRate:  expertEntity.AverageRate,
		TotalStudent: expertEntity.TotalStudent,
	}

	return course, expert, nil
}