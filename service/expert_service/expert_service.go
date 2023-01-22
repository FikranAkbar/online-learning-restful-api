package expert_service

import (
	"context"
	"online-learning-restful-api/model/web/expert"
)

type ExpertService interface {
	GetExpertDetailById(ctx context.Context, expertId uint) expert.DetailExpertResponse
}
