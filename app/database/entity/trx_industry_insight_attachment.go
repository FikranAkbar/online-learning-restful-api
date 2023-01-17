package entity

import (
	"time"
)

type TrxIndustryInsightAttachment struct {
	IndustryInsightId   uint `gorm:"column:industry_insight_id;primaryKey"`
	InsightAttachmentId uint `gorm:"column:insight_attachment_id;primaryKey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (TrxIndustryInsightAttachment) TableName() string {
	return "trx_industry_insight_attachment"
}
