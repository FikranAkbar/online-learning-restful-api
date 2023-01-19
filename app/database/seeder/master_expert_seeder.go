package seeder

import (
	"database/sql"
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterExpertSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterExpertSeeder(cfg gormseeder.SeederConfiguration) *MasterExpertSeeder {
	return &MasterExpertSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitExpertData() []entity.MasterExpert {
	return []entity.MasterExpert{
		{
			ID:         1,
			Name:       "Ray Robinson",
			Profession: sql.NullString{},
			ProfileDescription: sql.NullString{
				String: "Saya adalah expert Ray Robinson di platform online learning",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "081942142341",
				Valid:  true,
			},
			Address: sql.NullString{
				String: "Perum. Argopuro A-3",
				Valid:  true,
			},
			Gender:    "Male",
			BirthDate: sql.NullTime{},
			Photo: sql.NullString{
				String: "https://tilomitra.com/wp-content/uploads/2014/08/avatar-cartoon.png",
				Valid:  true,
			},
			AverageRate:  0,
			TotalStudent: 0,
			Categories:   nil,
		},
		{
			ID:         2,
			Name:       "Ivana Mcintosh",
			Profession: sql.NullString{},
			ProfileDescription: sql.NullString{
				String: "Saya adalah expert Ivana Mcintosh di platform online learning",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "082314240958",
				Valid:  true,
			},
			Address: sql.NullString{
				String: "Jl. Sumatera K-11",
				Valid:  true,
			},
			Gender:    "Female",
			BirthDate: sql.NullTime{},
			Photo: sql.NullString{
				String: "https://thumbs.dreamstime.com/b/young-woman-avatar-cartoon-character-profile-picture-young-brunette-woman-short-hair-avatar-cartoon-character-vector-149728784.jpg",
				Valid:  true,
			},
			AverageRate:  0,
			TotalStudent: 0,
			Categories:   nil,
		},
		{
			ID:         3,
			Name:       "Brielle Hutchinson",
			Profession: sql.NullString{},
			ProfileDescription: sql.NullString{
				String: "Saya adalah expert Brielle Hutchinson di platform online learning",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "081942142341",
				Valid:  true,
			},
			Address: sql.NullString{
				String: "Jl. Diponegoro 34",
				Valid:  true,
			},
			Gender:    "Female",
			BirthDate: sql.NullTime{},
			Photo: sql.NullString{
				String: "https://thumbs.dreamstime.com/b/young-woman-avatar-cartoon-character-profile-picture-young-brunette-woman-short-hair-avatar-cartoon-character-vector-149728784.jpg",
				Valid:  true,
			},
			AverageRate:  0,
			TotalStudent: 0,
			Categories:   nil,
		},
		{
			ID:         4,
			Name:       "Naomi Fadel",
			Profession: sql.NullString{},
			ProfileDescription: sql.NullString{
				String: "Saya adalah expert Naomi Fadel di platform online learning",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "084592482572",
				Valid:  true,
			},
			Address: sql.NullString{
				String: "Jl. Dr. Soetomo 2222",
				Valid:  true,
			},
			Gender:    "Male",
			BirthDate: sql.NullTime{},
			Photo: sql.NullString{
				String: "https://thumbs.dreamstime.com/b/girl-avatar-cartoon-stock-vector-image-cute-beautiful-eyes-93364804.jpg",
				Valid:  true,
			},
			AverageRate:  0,
			TotalStudent: 0,
			Categories:   nil,
		},
	}
}

func (s *MasterExpertSeeder) Seed(db *gorm.DB) error {
	experts := InitExpertData()

	return db.CreateInBatches(experts, len(experts)).Error
}

func (s *MasterExpertSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterExpert{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
