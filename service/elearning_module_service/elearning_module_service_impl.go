package elearning_module_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/elearning_module"
	"online-learning-restful-api/repository/elearning_module_repository"
)

type ElearningModuleServiceImpl struct {
	elearning_module_repository.ElearningModuleRepository
	*gorm.DB
}

func NewElearningModuleServiceImpl(elearningModuleRepository elearning_module_repository.ElearningModuleRepository, DB *gorm.DB) *ElearningModuleServiceImpl {
	return &ElearningModuleServiceImpl{ElearningModuleRepository: elearningModuleRepository, DB: DB}
}

func (service *ElearningModuleServiceImpl) GetOverviewElearningModulesByCourseId(ctx context.Context, courseId uint) []elearning_module.OverviewElearningModuleResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	elearningModules, err := service.ElearningModuleRepository.GetOverviewElearningModulesByCourseId(ctx, tx, courseId)
	helper.PanicIfError(err)

	var overviewElearningModulesResponse []elearning_module.OverviewElearningModuleResponse
	for _, elearningModule := range elearningModules {
		overviewElearningModulesResponse = append(overviewElearningModulesResponse, elearning_module.OverviewElearningModuleResponse{
			CourseId: elearningModule.CourseId,
			ModuleId: elearningModule.ID,
			Title:    elearningModule.ModuleTitle,
			Overview: elearningModule.ModuleOverview,
		})
	}

	return overviewElearningModulesResponse
}

