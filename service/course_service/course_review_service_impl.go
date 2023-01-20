package course_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/repository/course_repository"
)

type CourseReviewServiceImpl struct {
	course_repository.CourseRepository
	*gorm.DB
}

func NewCourseReviewServiceImpl(repository course_repository.CourseRepository, DB *gorm.DB) *CourseReviewServiceImpl {
	return &CourseReviewServiceImpl{CourseRepository: repository, DB: DB}
}

func (service *CourseReviewServiceImpl) GetCourseReviewsByCourseId(ctx context.Context, courseId uint) []course.ReviewCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courseReviews, err := service.CourseRepository.GetCourseReviewsByCourseId(ctx, tx, courseId)
	helper.PanicIfError(err)

	var reviewCourseResponses []course.ReviewCourseResponse
	for _, courseReview := range courseReviews {
		reviewCourseResponses = append(reviewCourseResponses, course.ReviewCourseResponse{
			UserId:     courseReview.UserId,
			UserName:   courseReview.UserName,
			ReviewDesc: courseReview.ReviewDesc,
			ReviewRate: courseReview.ReviewRate,
		})
	}

	return reviewCourseResponses
}
