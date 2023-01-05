package seeders

import (
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/core"
	"online-learning-restful-api/model/domain"
	"strconv"
	"time"
)

type MasterAccountSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewMasterAccountSeeder(cfg gorm_seeder.SeederConfiguration) *MasterAccountSeeder {
	return &MasterAccountSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (s *MasterAccountSeeder) Seed(db *gorm.DB) error {
	var accounts []domain.MasterAccount
	for i := 0; i < s.Configuration.Rows; i++ {
		account := domain.MasterAccount{
			DomainModel: core.DomainModel{
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			Email:    "foo" + strconv.Itoa(i) + "@bar.gov",
			Password: "password-hash-string" + strconv.Itoa(i),
			Role:     0,
		}
		accounts = append(accounts, account)
	}

	return db.CreateInBatches(accounts, s.Configuration.Rows).Error
}

func (s *MasterAccountSeeder) Clear(db *gorm.DB) error {
	entity := domain.MasterAccount{}
	return s.SeederAbstract.Truncate(db, entity.TableName())
}
