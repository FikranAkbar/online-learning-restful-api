package seeder

import (
	"database/sql"
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/model/entity"
	"time"
)

type MasterUserSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterUserSeeder(cfg gormseeder.SeederConfiguration) *MasterUserSeeder {
	return &MasterUserSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitUserData() []entity.MasterUser {
	return []entity.MasterUser{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "Molly Potts",
			Phone:     sql.NullString{String: "0811234567890"},
			Gender:    "Female",
			BirthDate: sql.NullTime{Time: time.Date(2007, 10, 11, 0, 0, 0, 0, time.UTC)},
			PhotoURL:  sql.NullString{String: "https://thumbs.dreamstime.com/b/young-woman-avatar-cartoon-character-profile-picture-young-brunette-woman-short-hair-avatar-cartoon-character-vector-149728784.jpg"},
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "Jack Guthrie",
			Phone:     sql.NullString{String: "081234567891"},
			Gender:    "Male",
			BirthDate: sql.NullTime{Time: time.Date(2005, 07, 04, 0, 0, 0, 0, time.UTC)},
			PhotoURL:  sql.NullString{String: "https://image.winudf.com/v2/image1/Y29tLm1uaWRlbmMuYXZ0YXJtYWtlcl9zY3JlZW5fNF8xNTU0NjQ5MzE4XzAwMw/screen-4.jpg?fakeurl=1&type=.jpg"},
		},
		{
			Model:     gorm.Model{ID: 3},
			Name:      "Finn Rhodes",
			Phone:     sql.NullString{String: "0811234567893"},
			Gender:    "Male",
			BirthDate: sql.NullTime{Time: time.Date(2011, 2, 8, 0, 0, 0, 0, time.UTC)},
			PhotoURL:  sql.NullString{String: "https://tilomitra.com/wp-content/uploads/2014/08/avatar-cartoon.png"},
		},
		{
			Model:     gorm.Model{ID: 4},
			Name:      "Clare Levy",
			Phone:     sql.NullString{String: "0811234567894"},
			Gender:    "Female",
			BirthDate: sql.NullTime{Time: time.Date(2008, 5, 14, 0, 0, 0, 0, time.UTC)},
			PhotoURL:  sql.NullString{String: "https://thumbs.dreamstime.com/b/girl-avatar-cartoon-stock-vector-image-cute-beautiful-eyes-93364804.jpg"},
		},
	}
}

func (s *MasterUserSeeder) Seed(db *gorm.DB) error {
	users := InitUserData()

	return db.CreateInBatches(users, len(users)).Error
}

func (s *MasterUserSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterUser{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
