package elearning_module_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/model/domain"
)

type ElearningModuleRepositoryImpl struct {
}

func NewElearningModuleRepositoryImpl() *ElearningModuleRepositoryImpl {
	return &ElearningModuleRepositoryImpl{}
}

func (repository *ElearningModuleRepositoryImpl) GetOverviewElearningModulesByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.ElearningModule, error) {
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		Preload("Modules").
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) || !courseEntity.IsPublished {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return []domain.ElearningModule{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return []domain.ElearningModule{}, err
	}

	moduleEntities := courseEntity.Modules
	err = db.WithContext(ctx).
		Where("course_id = ?", courseId).
		Preload("Sequence").
		Find(&moduleEntities).Error
	if err != nil {
		return []domain.ElearningModule{}, nil
	}

	var modules []domain.ElearningModule
	for _, moduleEntity := range moduleEntities {
		if !moduleEntity.IsPublished {
			continue
		}
		modules = append(modules, domain.ElearningModule{
			ID:             moduleEntity.ID,
			CourseId:       moduleEntity.CourseId,
			ModuleTitle:    moduleEntity.ModuleTitle,
			ModuleOverview: moduleEntity.ModuleOverview.String,
			IsPublished:    moduleEntity.IsPublished,
			SequenceNumber: moduleEntity.Sequence.OrderNumber,
		})
	}

	return modules, nil
}
