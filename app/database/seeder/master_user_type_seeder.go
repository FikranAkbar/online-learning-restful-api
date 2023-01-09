package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/model/entity"
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
			Model:    gorm.Model{ID: 1},
			UserType: "admin",
		},
		{
			Model:    gorm.Model{ID: 2},
			UserType: "user",
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
