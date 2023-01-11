package user_repository

import (
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(ctx context.Context, tx *gorm.Tx)
}
