package elearning_module_service

import (
	"context"
	"online-learning-restful-api/model/web/elearning_module"
)

type ElearningModuleService interface {
	GetOverviewElearningModulesByCourseId(ctx context.Context, courseId uint) []elearning_module.OverviewElearningModuleResponse
}
