package account_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type AccountRepository interface {
	FindUserByEmail(ctx context.Context, db *gorm.DB, email string) (*domain.Account, error)
}
