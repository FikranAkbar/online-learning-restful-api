package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterProvinceSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterProvinceSeeder(cfg gormseeder.SeederConfiguration) *MasterProvinceSeeder {
	return &MasterProvinceSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitProvinceData() []entity.MasterProvince {
	provinceStrings := []string{
		"ACEH", "SUMATERA UTARA", "SUMATERA BARAT", "RIAU", "JAMBI", "SUMATERA SELATAN", "BENGKULU", "LAMPUNG", "KEPULAUAN BANGKA BELITUNG", "KEPULAUAN RIAU", "DAERAH KHUSUS IBUKOTA JAKARTA", "JAWA BARAT", "JAWA TENGAH", "DAERAH ISTIMEWA YOGYAKARTA", "JAWA TIMUR", "BANTEN", "BALI", "NUSA TENGGARA BARAT", "NUSA TENGGARA TIMUR", "KALIMANTAN BARAT", "KALIMANTAN TENGAH", "KALIMANTAN SELATAN", "KALIMANTAN TIMUR", "KALIMANTAN UTARA", "SULAWESI UTARA", "SULAWESI TENGAH", "SULAWESI SELATAN", "SULAWESI TENGGARA", "GORONTALO", "SULAWESI BARAT", "MALUKU", "MALUKU UTARA", "PAPUA", "PAPUA BARAT",
	}

	var provinces []entity.MasterProvince

	for i, province := range provinceStrings {
		provinces = append(provinces, entity.MasterProvince{
			Model:        gorm.Model{ID: uint(i + 1)},
			ProvinceName: province,
		})
	}

	return provinces
}

func (s *MasterProvinceSeeder) Seed(db *gorm.DB) error {
	data := InitProvinceData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *MasterProvinceSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterProvince{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
