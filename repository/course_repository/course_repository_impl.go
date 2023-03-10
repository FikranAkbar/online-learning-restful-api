package course_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"strconv"
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
		fmt.Println(userEntity, "123")
		if err != nil && exception.CheckErrorContains(err, strings.ToLower(exception.NotFound)) {
			alreadyOwned = false
		} else if err != nil {
			helper.PanicIfError(err)
		} else {
			alreadyOwned = true
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

func (repository *CourseRepositoryImpl) GetUserCourseProgressionByCourseId(ctx context.Context, db *gorm.DB, courseId uint) (domain.UserCourse, error) {
	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	var userCourseEntity entity.TrxUserCourse
	err := db.WithContext(ctx).
		Where("user_id = ?", userTokenInfo.UserId).
		Where("course_id = ?", courseId).
		Preload("Course").
		First(&userCourseEntity).Error
	if exception.CheckErrorContains(err, exception.NotFound) {
		logError := "User not owned the access to course with id " + strconv.Itoa(int(courseId))
		return domain.UserCourse{}, exception.GenerateHTTPError(exception.Forbidden, logError)
	}

	userCourse := domain.UserCourse{
		UserId:               userCourseEntity.UserId,
		CourseId:             userCourseEntity.CourseId,
		LastUnlockedModule:   userCourseEntity.LastUnlockedModule,
		TotalWebinarAttended: userCourseEntity.TotalWebinarAttended,
		GraduatedAt:          userCourseEntity.GraduatedAt.Time,
	}
	return userCourse, nil
}

func (repository *CourseRepositoryImpl) GetCourseReviewsByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.CourseReview, error) {
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).Where("id = ?", courseId).Preload("Reviews").First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return []domain.CourseReview{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return []domain.CourseReview{}, err
	}

	courseReviewEntities := courseEntity.Reviews
	err = db.WithContext(ctx).Preload("User").Find(&courseReviewEntities).Error
	if err != nil {
		return []domain.CourseReview{}, err
	}

	var courseReviews []domain.CourseReview
	for _, courseReview := range courseReviewEntities {
		courseReviews = append(courseReviews, domain.CourseReview{
			Id:         courseReview.ID,
			CourseId:   courseReview.CourseId,
			UserId:     courseReview.UserId,
			UserName:   courseReview.User.Name,
			ReviewDesc: courseReview.ReviewDesc,
			ReviewRate: courseReview.ReviewRate,
		})
	}

	return courseReviews, nil
}

func (repository *CourseRepositoryImpl) GetComingSoonCourses(ctx context.Context, db *gorm.DB) ([]domain.Course, error) {
	var courseEntities []entity.MasterCourseComingSoon
	err := db.WithContext(ctx).Find(&courseEntities).Error
	if err != nil {
		return []domain.Course{}, err
	}

	var courses []domain.Course
	for _, courseEntity := range courseEntities {
		courses = append(courses, domain.Course{
			Id:       courseEntity.ID,
			Name:     courseEntity.Name,
			PhotoUrl: courseEntity.Cover.String,
		})
	}

	return courses, nil
}

func (repository *CourseRepositoryImpl) GetCourseSummary(ctx context.Context, db *gorm.DB, courseId uint) (domain.CourseSummary, error) {
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		Preload("Summary").
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return domain.CourseSummary{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.CourseSummary{}, err
	}

	courseSummaryEntity := courseEntity.Summary

	if courseSummaryEntity.ID == 0 {
		logError := fmt.Sprintf("Summary of the course with id %v not found", courseId)
		return domain.CourseSummary{}, exception.GenerateHTTPError(exception.NotFound, logError)
	}

	return domain.CourseSummary{
		ID:       courseSummaryEntity.ID,
		CourseId: courseSummaryEntity.CourseId,
		DocURL:   courseSummaryEntity.DocURL,
	}, nil
}

func (repository *CourseRepositoryImpl) CreateNewCourseOrder(ctx context.Context, db *gorm.DB, courseId uint) (domain.OrderCourse, error) {
	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	courseOrder := entity.TrxUserCart{
		UserId:   userTokenInfo.UserId,
		CourseId: courseId,
	}

	err := db.WithContext(ctx).
		Create(&courseOrder).Error
	if err != nil {
		fmt.Println(err)
		return domain.OrderCourse{}, err
	}

	return domain.OrderCourse{
		UserId:   courseOrder.UserId,
		CourseId: courseOrder.CourseId,
	}, nil
}

func (repository *CourseRepositoryImpl) DeleteCourseOrder(ctx context.Context, db *gorm.DB, courseId uint) (domain.OrderCourse, error) {
	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	courseOrder := entity.TrxUserCart{
		UserId:   userTokenInfo.UserId,
		CourseId: courseId,
	}

	err := db.WithContext(ctx).
		Delete(&courseOrder).Error
	if err != nil {
		fmt.Println(err)
		return domain.OrderCourse{}, err
	}

	return domain.OrderCourse{
		UserId:   courseOrder.UserId,
		CourseId: courseOrder.CourseId,
	}, nil
}
