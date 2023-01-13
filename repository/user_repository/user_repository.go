package user_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type UserRepository interface {
	FindUserById(ctx context.Context, tx *gorm.Tx, userId int) entity.MasterUser
}
