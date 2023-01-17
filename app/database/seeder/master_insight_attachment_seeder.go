package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterInsightAttachmentSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterInsightAttachmentSeeder(cfg gormseeder.SeederConfiguration) *MasterInsightAttachmentSeeder {
	return &MasterInsightAttachmentSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitInsightAttachmentData() []entity.MasterInsightAttachment {
	return []entity.MasterInsightAttachment{
		{
			ID:             1,
			AttachmentName: "Attachment 1",
		},
		{
			ID:             2,
			AttachmentName: "Attachment 2",
		},
		{
			ID:             3,
			AttachmentName: "Attachment 3",
		},
		{
			ID:             4,
			AttachmentName: "Attachment 4",
		},
		{
			ID:             5,
			AttachmentName: "Attachment 5",
		},
	}
}

func (s *MasterInsightAttachmentSeeder) Seed(db *gorm.DB) error {
	insightAttachments := InitInsightAttachmentData()

	return db.CreateInBatches(insightAttachments, len(insightAttachments)).Error
}

func (s *MasterInsightAttachmentSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterInsightAttachment{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
