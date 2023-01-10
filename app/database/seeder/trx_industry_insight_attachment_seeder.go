package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxIndustryInsightAttachmentSeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxIndustryInsightAttachmentSeeder(cfg gormseeder.SeederConfiguration) *TrxIndustryInsightAttachmentSeeder {
	return &TrxIndustryInsightAttachmentSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitIndustryInsightAttachmentRelationalData() []entity.TrxIndustryInsightAttachment {
	return []entity.TrxIndustryInsightAttachment{
		{
			IndustryInsightId:   1,
			InsightAttachmentId: 1,
		},
		{
			IndustryInsightId:   2,
			InsightAttachmentId: 2,
		},
		{
			IndustryInsightId:   3,
			InsightAttachmentId: 3,
		},
		{
			IndustryInsightId:   4,
			InsightAttachmentId: 4,
		},
		{
			IndustryInsightId:   4,
			InsightAttachmentId: 5,
		},
	}
}

func (s *TrxIndustryInsightAttachmentSeeder) Seed(db *gorm.DB) error {
	data := InitIndustryInsightAttachmentRelationalData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *TrxIndustryInsightAttachmentSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxIndustryInsightAttachment{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
