package entity

import "online-learning-restful-api/core"

type MasterInsightAttachment struct {
	core.EntityModel `gorm:"embedded"`
	AttachmentName   string `gorm:"type:varchar(100);not null"`
	DocURL           string `gorm:"not null"`
}

func (MasterInsightAttachment) TableName() string {
	return "master_insight_attachment"
}
