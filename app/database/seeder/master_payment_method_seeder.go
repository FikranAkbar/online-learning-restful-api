package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterPaymentMethodSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterPaymentMethodSeeder(cfg gormseeder.SeederConfiguration) *MasterPaymentMethodSeeder {
	return &MasterPaymentMethodSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitPaymentMethodData() entity.MasterPaymentMethod {
	return entity.MasterPaymentMethod{
		ID:   1,
		Name: "QRIS",
	}
}

func (seeder *MasterPaymentMethodSeeder) Seed(db *gorm.DB) error {
	data := InitPaymentMethodData()
	return db.Create(&data).Error
}

func (seeder *MasterPaymentMethodSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterPaymentMethod{}
	return seeder.SeederAbstract.Delete(db, ent.TableName())
}
