package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type MasterIndustryInsight struct {
	gorm.Model   `gorm:"embedded"`
	UserAuthorId uint                      `gorm:"column:user_author_id;not null"`
	MasterUser   MasterUser                `gorm:"foreignKey:UserAuthorId;joinForeignKey:ID"`
	TitleInsight string                    `gorm:"column:title_insight;type:varchar(100);not null"`
	CoverInsight sql.NullString            `gorm:"column:cover_insight"`
	BodyContent  sql.NullString            `gorm:"column:body_content"`
	IsPublished  bool                      `gorm:"column:is_published;default:false"`
	Attachments  []MasterInsightAttachment `gorm:"many2many:trx_industry_insight_attachment;foreignKey:ID;joinForeignKey:IndustryInsightId;references:ID;joinForeignKey:InsightAttachmentId"`
}

func (MasterIndustryInsight) TableName() string {
	return "master_industry_insight"
}
