package course_service

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/model/web/expert"
	"online-learning-restful-api/repository/course_repository"
)

type CourseDetailServiceImpl struct {
	course_repository.CourseRepository
	*gorm.DB
}

func NewCourseDetailServiceImpl(courseRepository course_repository.CourseRepository, DB *gorm.DB) *CourseDetailServiceImpl {
	return &CourseDetailServiceImpl{CourseRepository: courseRepository, DB: DB}
}

func (service *CourseDetailServiceImpl) GetCoursesByKeyword(ctx context.Context, keyword string) []course.ShortCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courses, err := service.CourseRepository.GetCoursesByKeyword(ctx, tx, keyword)
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

func (service *CourseDetailServiceImpl) GetDetailCourseByCourseId(ctx context.Context, courseId uint, userId *uint) course.DetailCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courseDetail, expertDetail, err := service.CourseRepository.GetCourseDetailByCourseId(ctx, tx, courseId, userId)
	helper.PanicIfError(err)

	contents := fmt.Sprintf("%vx e-learning modules and %vx webinar sessions", courseDetail.ModulesCount, courseDetail.WebinarsCount)

	detailCourseResponse := course.DetailCourseResponse{
		Id:                 courseDetail.Id,
		Name:               courseDetail.Name,
		Description:        courseDetail.Description,
		PhotoUrl:           courseDetail.PhotoUrl,
		Contents:           contents,
		AverageRate:        courseDetail.AverageRate,
		CurrentParticipant: courseDetail.CurrentParticipant,
		MaximumParticipant: courseDetail.MaximumParticipant,
		AlreadyOwned:       courseDetail.AlreadyOwned,
		Expert: expert.ShortExpertResponse{
			ID:    expertDetail.Id,
			Name:  expertDetail.Name,
			Photo: expertDetail.Photo,
		},
	}

	return detailCourseResponse
}

func (service *CourseDetailServiceImpl) GetUserCourseProgressionByCourseId(ctx context.Context, courseId uint) course.UserCourseProgressionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	userCourseProgression, err := service.CourseRepository.GetUserCourseProgressionByCourseId(ctx, tx, courseId)
	helper.PanicIfError(err)

	return course.UserCourseProgressionResponse{
		UserId:             userCourseProgression.UserId,
		CourseId:           userCourseProgression.CourseId,
		LastUnlockedModule: userCourseProgression.LastUnlockedModule,
		GraduatedAt:        userCourseProgression.GraduatedAt,
	}
}
