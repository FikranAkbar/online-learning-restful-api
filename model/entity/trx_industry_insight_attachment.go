package entity

import "online-learning-restful-api/core"

type TrxIndustryInsightAttachment struct {
	core.EntityModel        `gorm:"embedded"`
	IndustryInsightId       uint                    `gorm:"column:industry_insight_id;not null"`
	MasterIndustryInsight   MasterIndustryInsight   `gorm:"foreignKey:IndustryInsightId"`
	InsightAttachmentId     uint                    `gorm:"column:insight_attachment_id;not null"`
	MasterInsightAttachment MasterInsightAttachment `gorm:"foreignKey:InsightAttachmentId"`
}

func (TrxIndustryInsightAttachment) TableName() string {
	return "trx_industry_insight_attachment"
}
