package domain

import "online-learning-restful-api/core"

type TrxIndustryInsightAttachment struct {
	core.DomainModel
	IndustryInsightId   uint `gorm:"not null"`
	InsightAttachmentId uint `gorm:"not null"`
}

func (TrxIndustryInsightAttachment) TableName() string {
	return "trx_industry_insight_attachment"
}
