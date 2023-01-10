package seeder

import (
	"database/sql"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterCourseComingSoonSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewMasterCourseComingSoonSeeder(cfg gorm_seeder.SeederConfiguration) *MasterCourseComingSoonSeeder {
	return &MasterCourseComingSoonSeeder{gorm_seeder.SeederAbstract{Configuration: cfg}}
}

func InitComingSoonCourseData() []entity.MasterCourseComingSoon {
	return []entity.MasterCourseComingSoon{
		{
			Model:       gorm.Model{ID: 1},
			Name:        "Entrepreneurship",
			Description: sql.NullString{String: "Course Enterpreneurship sedang dalam pengerjaan dan akan hadir segera."},
			Cover:       sql.NullString{String: "https://www.alphajwc.com/wp-content/uploads/2020/02/pengusaha-muda-indonesia.jpg"},
			IsPublished: true,
		},
		{
			Model:       gorm.Model{ID: 2},
			Name:        "Germany Language",
			Description: sql.NullString{String: "Course Bahasa Jerman sedang dalam pengerjaan dan akan hadir segera."},
			Cover:       sql.NullString{String: "https://www.banksinarmas.com/biasakansekarang/wp-content/uploads/2018/11/image5.png"},
			IsPublished: true,
		},
		{
			Model:       gorm.Model{ID: 3},
			Name:        "Leadership",
			Description: sql.NullString{String: "Course Leadership sedang dalam pengerjaan dan akan hadir segera."},
			Cover:       sql.NullString{String: "https://cdn.futuready.com/artikel/media/2017/08/26105544/37257994_l.jpg"},
			IsPublished: true,
		},
		{
			Model:       gorm.Model{ID: 4},
			Name:        "Programming",
			Description: sql.NullString{String: "Course Programming sedang dalam pengerjaan dan akan hadir segera."},
			Cover:       sql.NullString{String: "https://res.cloudinary.com/sisternet-co-id/image/upload/q_auto:best,f_auto/article/vynfvmlwmthncwkn5enl.jpg"},
			IsPublished: true,
		},
	}
}

func (s *MasterCourseComingSoonSeeder) Seed(db *gorm.DB) error {
	comingSoonCourses := InitComingSoonCourseData()

	return db.CreateInBatches(comingSoonCourses, len(comingSoonCourses)).Error
}

func (s *MasterCourseComingSoonSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterCourseComingSoon{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
