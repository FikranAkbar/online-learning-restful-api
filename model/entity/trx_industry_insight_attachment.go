package entity

import (
	"gorm.io/gorm"
)

type TrxIndustryInsightAttachment struct {
	gorm.Model              `gorm:"embedded"`
	IndustryInsightId       uint                    `gorm:"column:industry_insight_id;not null"`
	MasterIndustryInsight   MasterIndustryInsight   `gorm:"foreignKey:IndustryInsightId"`
	InsightAttachmentId     uint                    `gorm:"column:insight_attachment_id;not null"`
	MasterInsightAttachment MasterInsightAttachment `gorm:"foreignKey:InsightAttachmentId"`
}

func (TrxIndustryInsightAttachment) TableName() string {
	return "trx_industry_insight_attachment"
}
