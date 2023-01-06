package entity

import "online-learning-restful-api/core"

type TrxIndustryInsightAttachment struct {
	core.EntityModel    `gorm:"embedded"`
	IndustryInsightId   uint `gorm:"not null"`
	InsightAttachmentId uint `gorm:"not null"`
}

func (TrxIndustryInsightAttachment) TableName() string {
	return "trx_industry_insight_attachment"
}
