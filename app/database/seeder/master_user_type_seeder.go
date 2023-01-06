package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/core"
	"online-learning-restful-api/model/entity"
)

type MasterUserTypeSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterUserTypeSeeder(cfg gormseeder.SeederConfiguration) *MasterUserTypeSeeder {
	return &MasterUserTypeSeeder{gormseeder.NewSeederAbstract(cfg)}
}

func (s *MasterUserTypeSeeder) Seed(db *gorm.DB) error {
	adminRole := entity.MasterUserType{
		EntityModel: core.EntityModel{ID: 1},
		UserType:    "admin",
	}
	userRole := entity.MasterUserType{
		EntityModel: core.EntityModel{ID: 2},
		UserType:    "user",
	}

	userTypes := []entity.MasterUserType{adminRole, userRole}

	return db.CreateInBatches(userTypes, len(userTypes)).Error
}

func (s *MasterUserTypeSeeder) Clear(db *gorm.DB) error {
	ent := entity.MasterUserType{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
