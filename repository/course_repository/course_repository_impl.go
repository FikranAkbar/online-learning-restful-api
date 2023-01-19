package course_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/model/domain"
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
