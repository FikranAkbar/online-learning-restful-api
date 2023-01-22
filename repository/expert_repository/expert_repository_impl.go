package expert_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/model/domain"
)

type ExpertRepositoryImpl struct {
}

func NewExpertRepositoryImpl() *ExpertRepositoryImpl {
	return &ExpertRepositoryImpl{}
}

func (repository *ExpertRepositoryImpl) GetExpertDetailById(ctx context.Context, db *gorm.DB, expertId uint) (domain.Expert, error) {
	var expertEntity entity.MasterExpert
	err := db.WithContext(ctx).
		Where("id = ?", expertId).
		First(&expertEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("Expert with id %v not found", expertId)
		return domain.Expert{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.Expert{}, err
	}

	var courseEntities []entity.MasterCourse
	err = db.WithContext(ctx).
		Where("expert_id = ?", expertId).
		Find(&courseEntities).Error
	if err != nil {
		return domain.Expert{}, err
	}

	return domain.Expert{
		Id:           expertEntity.ID,
		Name:         expertEntity.Name,
		Profession:   expertEntity.Profession.String,
		Phone:        expertEntity.Phone.String,
		Address:      expertEntity.Address.String,
		Gender:       expertEntity.Gender,
		Photo:        expertEntity.Photo.String,
		BirthDate:    expertEntity.BirthDate.Time,
		AverageRate:  expertEntity.AverageRate,
		TotalStudent: expertEntity.TotalStudent,
		ReviewsCount: len(courseEntities),
	}, nil
}
