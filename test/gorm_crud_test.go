package test

import (
	"fmt"
	"gorm.io/gorm/clause"
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/entity"
	"testing"
)

func TestFirstData(t *testing.T) {
	db := database.NewDB()
	var account *entity.MasterAccount
	db.Preload("MasterUserType").First(&account)
	fmt.Print(helper.StringModel(*account))
}

func TestFindDataAndRelationshipData(t *testing.T) {
	db := database.NewDB()
	var accounts []entity.MasterAccount
	db.Preload(clause.Associations).Find(&accounts)

	for _, val := range accounts {
		fmt.Printf("Name %v Role %v\n", val.Email, val.MasterUserType.Name)
	}
}

func TestFindDataManyToMany(t *testing.T) {
	db := database.NewDB()
	var users []entity.MasterUser
	db.Model(&entity.MasterUser{}).Preload(clause.Associations).Find(&users)

	for _, user := range users {
		fmt.Printf("My name is %v and my course is", user.Name)

		for _, course := range user.Courses {
			fmt.Printf(" %v", course.Name)
		}

		fmt.Println()
	}
}
