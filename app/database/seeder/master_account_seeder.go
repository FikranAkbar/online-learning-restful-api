package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterAccountSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterAccountSeeder(cfg gormseeder.SeederConfiguration) *MasterAccountSeeder {
	return &MasterAccountSeeder{gormseeder.NewSeederAbstract(cfg)}
}

func InitAccountData() []entity.MasterAccount {
	return []entity.MasterAccount{
		{
			ID:       1,
			Email:    "mollypotts@gmail.com",
			Password: "$2a$10$BQEJrr7oDIBK3QxrmK16e.gOkjpy4h0O3RLW35SUVRGGyoAq7EO3m",
			Role:     2,
		},
		{
			ID:       2,
			Email:    "jackguthrie@gmail.com",
			Password: "$2a$10$BQEJrr7oDIBK3QxrmK16e.gOkjpy4h0O3RLW35SUVRGGyoAq7EO3m",
			Role:     2,
		},
		{
			ID:       3,
			Email:    "finnrhodes@gmail.com",
			Password: "$2a$10$BQEJrr7oDIBK3QxrmK16e.gOkjpy4h0O3RLW35SUVRGGyoAq7EO3m",
			Role:     2,
		},
		{
			ID:       4,
			Email:    "claralevy@gmail.com",
			Password: "$2a$10$BQEJrr7oDIBK3QxrmK16e.gOkjpy4h0O3RLW35SUVRGGyoAq7EO3m",
			Role:     2,
		},
	}
}

func (s *MasterAccountSeeder) Seed(db *gorm.DB) error {
	accounts := InitAccountData()

	return db.CreateInBatches(accounts, len(accounts)).Error
}

func (s *MasterAccountSeeder) Clear(db *gorm.DB) error {
	ent := entity.MasterAccount{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
