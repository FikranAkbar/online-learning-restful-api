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
		fmt.Printf("Name %v Role %v\n", val.Email, val.MasterUserType.UserType)
	}
}

func TestFindDataManyToMany(t *testing.T) {
	db := database.NewDB()
	var experts []entity.MasterExpert
	db.Model(&entity.MasterExpert{}).Preload("Categories").Find(&experts)

	for _, expert := range experts {
		fmt.Printf("My name is %v I will handle", expert.Name)

		for _, category := range expert.Categories {
			fmt.Printf(" %v", category.CategoryName)
		}

		fmt.Println()
	}
}
