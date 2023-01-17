package seeder

import (
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterVideoSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewMasterVideoSeeder(cfg gorm_seeder.SeederConfiguration) *MasterVideoSeeder {
	return &MasterVideoSeeder{gorm_seeder.SeederAbstract{Configuration: cfg}}
}

func InitVideoData() []entity.MasterVideo {
	return []entity.MasterVideo{
		//#Video for Intro to Entrepreneurship
		{
			ID:          1,
			CourseId:    1,
			ModuleId:    1,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          2,
			CourseId:    1,
			ModuleId:    2,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          3,
			CourseId:    1,
			ModuleId:    3,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          4,
			CourseId:    1,
			ModuleId:    4,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Intermediate to Entrepreneurship
		{
			ID:          5,
			CourseId:    5,
			ModuleId:    5,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          6,
			CourseId:    5,
			ModuleId:    6,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          7,
			CourseId:    5,
			ModuleId:    7,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          8,
			CourseId:    5,
			ModuleId:    8,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Expert to Entrepreneurship
		{
			ID:          9,
			CourseId:    9,
			ModuleId:    9,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          10,
			CourseId:    9,
			ModuleId:    10,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          11,
			CourseId:    9,
			ModuleId:    11,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          12,
			CourseId:    9,
			ModuleId:    12,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Intro to Marketing
		{
			ID:          13,
			CourseId:    2,
			ModuleId:    13,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          14,
			CourseId:    2,
			ModuleId:    14,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          15,
			CourseId:    2,
			ModuleId:    15,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          16,
			CourseId:    2,
			ModuleId:    16,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Intermediate to Marketing
		{
			ID:          17,
			CourseId:    6,
			ModuleId:    17,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          18,
			CourseId:    6,
			ModuleId:    18,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          19,
			CourseId:    6,
			ModuleId:    19,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          20,
			CourseId:    6,
			ModuleId:    20,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Expert to Marketing
		{
			ID:          21,
			CourseId:    10,
			ModuleId:    21,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          22,
			CourseId:    10,
			ModuleId:    22,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          23,
			CourseId:    10,
			ModuleId:    23,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          24,
			CourseId:    10,
			ModuleId:    24,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Intro to English
		{
			ID:          25,
			CourseId:    3,
			ModuleId:    25,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          26,
			CourseId:    3,
			ModuleId:    26,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          27,
			CourseId:    3,
			ModuleId:    27,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          28,
			CourseId:    3,
			ModuleId:    28,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Intermediate to English
		{
			ID:          29,
			CourseId:    7,
			ModuleId:    29,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          30,
			CourseId:    7,
			ModuleId:    30,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          31,
			CourseId:    7,
			ModuleId:    31,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          32,
			CourseId:    7,
			ModuleId:    32,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Expert to English
		{
			ID:          33,
			CourseId:    11,
			ModuleId:    33,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          34,
			CourseId:    11,
			ModuleId:    34,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          35,
			CourseId:    11,
			ModuleId:    35,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          36,
			CourseId:    11,
			ModuleId:    36,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Intro to Creative
		{
			ID:          37,
			CourseId:    4,
			ModuleId:    37,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          38,
			CourseId:    4,
			ModuleId:    38,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          39,
			CourseId:    4,
			ModuleId:    39,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          40,
			CourseId:    4,
			ModuleId:    40,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Intermediate to Creative
		{
			ID:          41,
			CourseId:    8,
			ModuleId:    41,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          42,
			CourseId:    8,
			ModuleId:    42,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          43,
			CourseId:    8,
			ModuleId:    43,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          44,
			CourseId:    8,
			ModuleId:    44,
			VideoName:   "Video 4",
			IsPublished: true,
		},

		//#Video for Expert to English
		{
			ID:          45,
			CourseId:    12,
			ModuleId:    45,
			VideoName:   "Video 1",
			IsPublished: true,
		},
		{
			ID:          46,
			CourseId:    12,
			ModuleId:    46,
			VideoName:   "Video 2",
			IsPublished: true,
		},
		{
			ID:          47,
			CourseId:    12,
			ModuleId:    47,
			VideoName:   "Video 3",
			IsPublished: true,
		},
		{
			ID:          48,
			CourseId:    12,
			ModuleId:    48,
			VideoName:   "Video 4",
			IsPublished: true,
		},
	}
}

func (s *MasterVideoSeeder) Seed(db *gorm.DB) error {
	videos := InitVideoData()

	return db.CreateInBatches(videos, len(videos)).Error
}

func (s *MasterVideoSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterVideo{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
