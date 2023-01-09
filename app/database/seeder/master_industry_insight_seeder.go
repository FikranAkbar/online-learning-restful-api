package seeder

import (
	"database/sql"
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/model/entity"
)

type MasterIndustryInsightSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterIndustryInsightSeeder(cfg gormseeder.SeederConfiguration) *MasterIndustryInsightSeeder {
	return &MasterIndustryInsightSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitIndustryInsightData() []entity.MasterIndustryInsight {
	return []entity.MasterIndustryInsight{
		{
			UserAuthorId: 1,
			TitleInsight: "Insight Title 1",
			CoverInsight: sql.NullString{String: "https://images.unsplash.com/photo-1517732306149-e8f829eb588a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80"},
			BodyContent:  sql.NullString{String: "Ini adalah deskripsi insight title 1 Ini adalah deskripsi insight title 1 Ini adalah deskripsi insight title 1 Ini adalah deskripsi insight title 1 Ini adalah deskripsi insight title 1 Ini adalah deskripsi insight title 1 Ini adalah deskripsi insight title 1 Ini adalah deskripsi insight title 1 Ini adalah deskripsi insight title 1"},
			IsPublished:  true,
		},
		{
			UserAuthorId: 2,
			TitleInsight: "Insight Title 2",
			CoverInsight: sql.NullString{String: "https://images.unsplash.com/photo-1517732306149-e8f829eb588a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80"},
			BodyContent:  sql.NullString{String: "Ini adalah deskripsi insight title 2 Ini adalah deskripsi insight title 2 Ini adalah deskripsi insight title 2 Ini adalah deskripsi insight title 2 Ini adalah deskripsi insight title 2 Ini adalah deskripsi insight title 2 Ini adalah deskripsi insight title 2 Ini adalah deskripsi insight title 2 Ini adalah deskripsi insight title 2"},
			IsPublished:  true,
		},
		{
			UserAuthorId: 3,
			TitleInsight: "Insight Title 3",
			CoverInsight: sql.NullString{String: "https://images.unsplash.com/photo-1521517407911-565264e7d82d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80"},
			BodyContent:  sql.NullString{String: "Ini adalah deskripsi insight title 3 Ini adalah deskripsi insight title 3 Ini adalah deskripsi insight title 3 Ini adalah deskripsi insight title 3 Ini adalah deskripsi insight title 3 Ini adalah deskripsi insight title 3 Ini adalah deskripsi insight title 3 Ini adalah deskripsi insight title 3 Ini adalah deskripsi insight title 3"},
			IsPublished:  true,
		},
		{
			UserAuthorId: 4,
			TitleInsight: "Insight Title 4",
			CoverInsight: sql.NullString{String: "https://images.unsplash.com/photo-1508480408285-812f971170b3?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80"},
			BodyContent:  sql.NullString{String: "Ini adalah deskripsi insight title 4 Ini adalah deskripsi insight title 4 Ini adalah deskripsi insight title 4 Ini adalah deskripsi insight title 4 Ini adalah deskripsi insight title 4 Ini adalah deskripsi insight title 4 Ini adalah deskripsi insight title 4 Ini adalah deskripsi insight title 4 Ini adalah deskripsi insight title 4"},
			IsPublished:  true,
		},
	}
}

func (s *MasterIndustryInsightSeeder) Seed(db *gorm.DB) error {
	industryInsights := InitIndustryInsightData()

	return db.CreateInBatches(industryInsights, len(industryInsights)).Error
}

func (s *MasterIndustryInsightSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterIndustryInsight{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
