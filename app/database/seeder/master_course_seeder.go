package seeder

import (
	"database/sql"
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterCourseSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterCourseSeeder(cfg gormseeder.SeederConfiguration) *MasterCourseSeeder {
	return &MasterCourseSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitCourseData() []entity.MasterCourse {
	return []entity.MasterCourse{
		{
			Model:    gorm.Model{ID: 1},
			ExpertId: 1,
			StatusId: 2,
			Name:     "Intro to Entrepreneurship",
			Description: sql.NullString{
				String: "Deskripsi Intro to Entrepreneurship Deskripsi Intro to Entrepreneurship Deskripsi Intro to Entrepreneurship Deskripsi Intro to Entrepreneurship.",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1432888498266-38ffec3eaf0a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1174&q=80",
			},
			Price:              200000,
			TotalDuration:      24.5,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 2},
			ExpertId: 2,
			StatusId: 2,
			Name:     "Intro to Marketing",
			Description: sql.NullString{
				String: "Deskripsi Intro to Marketing Deskripsi Intro to Marketing Deskripsi Intro to Marketing Deskripsi Intro to Marketing Deskripsi Intro to Marketing.",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1434626881859-194d67b2b86f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1174&q=80",
			},
			Price:              300000,
			TotalDuration:      20.2,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 3},
			ExpertId: 3,
			StatusId: 2,
			Name:     "Intro to English",
			Description: sql.NullString{
				String: "Deskripsi Intro to English Deskripsi Intro to English Deskripsi Intro to English Deskripsi Intro to English",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1518463892881-d587bf2c296a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1228&q=80",
			},
			Price:              250000,
			TotalDuration:      25,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 4},
			ExpertId: 4,
			StatusId: 2,
			Name:     "Intro to Creative",
			Description: sql.NullString{
				String: "Deskripsi Intro to Creative Deskripsi Intro to Creative Deskripsi Intro to Creative Deskripsi Intro to Creative",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1628359355624-855775b5c9c4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80",
			},
			Price:              400000,
			TotalDuration:      21.5,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 5},
			ExpertId: 3,
			StatusId: 2,
			Name:     "Intermediate to Entrepreneurship",
			Description: sql.NullString{
				String: "Deskripsi Intermediate to Entrepreneurship Deskripsi Intermediate to Entrepreneurship Deskripsi Intermediate to Entrepreneurship",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1517245386807-bb43f82c33c4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80",
			},
			Price:              220000,
			TotalDuration:      19,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 6},
			ExpertId: 1,
			StatusId: 2,
			Name:     "Intermediate to Marketing",
			Description: sql.NullString{
				String: "Deskripsi Intermedia to Marketing Deskripsi Intermedia to Marketing Deskripsi Intermedia to Marketing Deskripsi Intermedia to Marketing",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1533750349088-cd871a92f312?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80",
			},
			Price:              350000,
			TotalDuration:      14.2,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 7},
			ExpertId: 4,
			StatusId: 2,
			Name:     "Intermediate to English",
			Description: sql.NullString{
				String: "Deskripsi Intermediate to English Deskripsi Intermediate to English Deskripsi Intermediate to English Deskripsi Intermediate to English",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1571260899304-425eee4c7efc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80",
			},
			Price:              160000,
			TotalDuration:      19.8,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 8},
			ExpertId: 4,
			StatusId: 2,
			Name:     "Intermediate to Creative",
			Description: sql.NullString{
				String: "Deskripsi Intermediate to Creative Deskripsi Intermediate to Creative Deskripsi Intermediate to Creative Deskripsi Intermediate to Creative",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1174&q=80",
			},
			Price:              130000,
			TotalDuration:      18.4,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 9},
			ExpertId: 1,
			StatusId: 2,
			Name:     "Expert to Entrepreneurship",
			Description: sql.NullString{
				String: "Deskripsi Expert to Entrepreneurship Deskripsi Expert to Entrepreneurship Deskripsi Expert to Entrepreneurship Deskripsi Expert to Entrepreneurship",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1533749871411-5e21e14bcc7d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1171&q=80",
			},
			Price:              490000,
			TotalDuration:      11.1,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 10},
			ExpertId: 2,
			StatusId: 2,
			Name:     "Expert to Marketing",
			Description: sql.NullString{
				String: "Deskripsi Expert to Marketing Deskripsi Expert to Marketing Deskripsi Expert to Marketing Deskripsi Expert to Marketing.",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1533749871411-5e21e14bcc7d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1171&q=80",
			},
			Price:              420000,
			TotalDuration:      13.3,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 11},
			ExpertId: 3,
			StatusId: 2,
			Name:     "Expert to English",
			Description: sql.NullString{
				String: "Deskripsi Expert to English Deskripsi Expert to English Deskripsi Expert to English.",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1567206163313-9e34c830557a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1171&q=80",
			},
			Price:              370000,
			TotalDuration:      17.1,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
		{
			Model:    gorm.Model{ID: 12},
			ExpertId: 4,
			StatusId: 2,
			Name:     "Expert to Creative",
			Description: sql.NullString{
				String: "Deskripsi Expert to Creative Deskripsi Expert to Creative Deskripsi Expert to Creative Deskripsi Expert to Creative Deskripsi Expert to Creative.",
			},
			PhotoURL: sql.NullString{
				String: "https://images.unsplash.com/photo-1501084817091-a4f3d1d19e07?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTN8fGNyZWF0aXZlfGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=600&q=60",
			},
			Price:              540000,
			TotalDuration:      18.4,
			CurrentParticipant: 0,
			MaximumParticipant: 30,
			IsPublished:        true,
		},
	}
}

func (s *MasterCourseSeeder) Seed(db *gorm.DB) error {
	courses := InitCourseData()

	return db.CreateInBatches(courses, len(courses)).Error
}

func (s *MasterCourseSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterCourse{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
