package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/model/entity"
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
			Model:          gorm.Model{ID: 1},
			AttachmentName: "Attachment 1",
		},
		{
			Model:          gorm.Model{ID: 2},
			AttachmentName: "Attachment 2",
		},
		{
			Model:          gorm.Model{ID: 3},
			AttachmentName: "Attachment 3",
		},
		{
			Model:          gorm.Model{ID: 4},
			AttachmentName: "Attachment 4",
		},
		{
			Model:          gorm.Model{ID: 5},
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
