package entity

import "online-learning-restful-api/core"

type TrxIndustryInsightAttachment struct {
	core.EntityModel        `gorm:"embedded"`
	IndustryInsightId       uint                    `gorm:"not null"`
	MasterIndustryInsight   MasterIndustryInsight   `gorm:"foreignKey:IndustryInsightId"`
	InsightAttachmentId     uint                    `gorm:"not null"`
	MasterInsightAttachment MasterInsightAttachment `gorm:"foreignKey:InsightAttachmentId"`
}

func (TrxIndustryInsightAttachment) TableName() string {
	return "trx_industry_insight_attachment"
}
