package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterUserTypeSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterUserTypeSeeder(cfg gormseeder.SeederConfiguration) *MasterUserTypeSeeder {
	return &MasterUserTypeSeeder{gormseeder.NewSeederAbstract(cfg)}
}

func InitUserTypeData() []entity.MasterUserType {
	return []entity.MasterUserType{
		{
			ID:   1,
			Name: "admin",
		},
		{
			ID:   2,
			Name: "user",
		},
	}
}

func (s *MasterUserTypeSeeder) Seed(db *gorm.DB) error {
	userTypes := InitUserTypeData()

	return db.CreateInBatches(userTypes, len(userTypes)).Error
}

func (s *MasterUserTypeSeeder) Clear(db *gorm.DB) error {
	ent := entity.MasterUserType{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
