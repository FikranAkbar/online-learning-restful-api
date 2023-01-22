package elearning_module_service

import (
	"context"
	"online-learning-restful-api/model/web/elearning_module"
	"online-learning-restful-api/model/web/video"
)

type ElearningModuleService interface {
	GetOverviewElearningModulesByCourseId(ctx context.Context, courseId uint) []elearning_module.OverviewElearningModuleResponse
	GetDetailElearningModuleByElearningModuleId(ctx context.Context, courseId uint, elearningModuleId uint) elearning_module.DetailElearningModuleResponse
	SaveVideoProgressionInModule(ctx context.Context, courseId uint, elearningModuleId uint, request video.UserVideoProgressionRequest) video.UserVideoProgressionResponse
}
