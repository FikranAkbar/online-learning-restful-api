package seeder

import (
	"database/sql"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterWebinarSessionSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewMasterWebinarSessionSeeder(cfg gorm_seeder.SeederConfiguration) *MasterWebinarSessionSeeder {
	return &MasterWebinarSessionSeeder{gorm_seeder.SeederAbstract{Configuration: cfg}}
}

func InitWebinarSessionData() []entity.MasterWebinarSession {
	var scheduleDay1 uint = 4
	scheduleDate1 := "2022-10-05"
	timeStart1 := "09:00"
	timeFinish1 := "11:00"

	var scheduleDay2 uint = 7
	scheduleDate2 := "2022-10-08"
	timeStart2 := "10:00"
	timeFinish2 := "12:00"

	return []entity.MasterWebinarSession{
		//#Webinar Session Intro to Entrepreneurship
		{
			Model:        gorm.Model{ID: 1},
			CourseId:     1,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 2},
			CourseId:     1,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 3},
			CourseId:     1,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 4},
			CourseId:     1,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Intermediate to Entrepreneurship
		{
			Model:        gorm.Model{ID: 5},
			CourseId:     5,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 6},
			CourseId:     5,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 7},
			CourseId:     5,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 8},
			CourseId:     5,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Expert to Entrepreneurship
		{
			Model:        gorm.Model{ID: 9},
			CourseId:     9,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 10},
			CourseId:     9,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 11},
			CourseId:     9,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 12},
			CourseId:     9,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Intro to Marketing
		{
			Model:        gorm.Model{ID: 13},
			CourseId:     2,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 14},
			CourseId:     2,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 15},
			CourseId:     2,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 16},
			CourseId:     2,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Intermediate to Marketing
		{
			Model:        gorm.Model{ID: 17},
			CourseId:     6,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 18},
			CourseId:     6,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 19},
			CourseId:     6,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 20},
			CourseId:     6,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Expert to Marketing
		{
			Model:        gorm.Model{ID: 21},
			CourseId:     10,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 22},
			CourseId:     10,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 23},
			CourseId:     10,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 24},
			CourseId:     10,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Intro to English
		{
			Model:        gorm.Model{ID: 25},
			CourseId:     3,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 26},
			CourseId:     3,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 27},
			CourseId:     3,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 28},
			CourseId:     3,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Intermediate to English
		{
			Model:        gorm.Model{ID: 29},
			CourseId:     7,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 30},
			CourseId:     7,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 31},
			CourseId:     7,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 32},
			CourseId:     7,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Expert to English
		{
			Model:        gorm.Model{ID: 33},
			CourseId:     11,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 34},
			CourseId:     11,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 35},
			CourseId:     11,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 36},
			CourseId:     11,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Intro to Creative
		{
			Model:        gorm.Model{ID: 37},
			CourseId:     4,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 38},
			CourseId:     4,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 39},
			CourseId:     4,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 40},
			CourseId:     4,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Intermediate to Creative
		{
			Model:        gorm.Model{ID: 41},
			CourseId:     8,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 42},
			CourseId:     8,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 43},
			CourseId:     8,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 44},
			CourseId:     8,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#

		//#Webinar Session Expert to Creative
		{
			Model:        gorm.Model{ID: 45},
			CourseId:     12,
			Title:        "Webinar Title 1",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1 Deskripsi dari Webinar Title 1"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 46},
			CourseId:     12,
			Title:        "Webinar Title 2",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2 Deskripsi dari Webinar Title 2"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 47},
			CourseId:     12,
			Title:        "Webinar Title 3",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3 Deskripsi dari Webinar Title 3"},
			ScheduleDay:  scheduleDay2,
			ScheduleDate: scheduleDate2,
			TimeStart:    timeStart2,
			TimeFinish:   timeFinish2,
			IsPublished:  true,
		},
		{
			Model:        gorm.Model{ID: 48},
			CourseId:     12,
			Title:        "Webinar Title 4",
			Desc:         sql.NullString{String: "Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4 Deskripsi dari Webinar Title 4"},
			ScheduleDay:  scheduleDay1,
			ScheduleDate: scheduleDate1,
			TimeStart:    timeStart1,
			TimeFinish:   timeFinish1,
			IsPublished:  true,
		},
		//#
	}
}

func (s *MasterWebinarSessionSeeder) Seed(db *gorm.DB) error {
	webinarSessions := InitWebinarSessionData()

	return db.CreateInBatches(webinarSessions, len(webinarSessions)).Error
}

func (s *MasterWebinarSessionSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterWebinarSession{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
