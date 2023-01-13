package account_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type AccountRepository interface {
	FindUserByEmail(ctx context.Context, tx *gorm.Tx, email string) entity.MasterAccount
}
