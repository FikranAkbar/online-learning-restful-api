package entity

import (
	"database/sql"
	"online-learning-restful-api/core"
)

type MasterIndustryInsight struct {
	core.EntityModel `gorm:"embedded"`
	UserAuthorId     uint       `gorm:"not null"`
	MasterUser       MasterUser `gorm:"foreignKey:UserAuthorId"`
	TitleInsight     string     `gorm:"type:varchar(100);not null"`
	CoverInsight     sql.NullString
	BodyContent      sql.NullString
	IsPublished      bool `gorm:"default:false"`
}

func (MasterIndustryInsight) TableName() string {
	return "master_industry_insight"
}
