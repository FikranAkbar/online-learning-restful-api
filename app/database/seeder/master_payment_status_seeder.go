package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterPaymentStatusSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterPaymentStatusSeeder(cfg gormseeder.SeederConfiguration) *MasterPaymentStatusSeeder {
	return &MasterPaymentStatusSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitPaymentStatusData() []entity.MasterPaymentStatus {
	return []entity.MasterPaymentStatus{
		{
			ID:   1,
			Name: "completed",
		},
		{
			ID:   2,
			Name: "cancelled",
		},
		{
			ID:   3,
			Name: "waiting",
		},
	}
}

func (s *MasterPaymentStatusSeeder) Seed(db *gorm.DB) error {
	data := InitPaymentStatusData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *MasterPaymentStatusSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterPaymentStatus{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
