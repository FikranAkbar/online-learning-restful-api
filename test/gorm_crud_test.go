package test

import (
	"fmt"
	"gorm.io/gorm/clause"
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/model/entity"
	"testing"
)

func TestFindData(t *testing.T) {
	db := database.NewDB()
	var accounts *[]entity.MasterAccount
	db.Find(&accounts)

	for _, val := range *accounts {
		fmt.Println(val)
	}
}

func TestFindDataAndRelationshipData(t *testing.T) {
	db := database.NewDB()
	var accounts []entity.MasterAccount
	db.Preload(clause.Associations).Find(&accounts)

	for _, val := range accounts {
		fmt.Println()
		fmt.Printf("Name %v Role %v\n", val.Email, val.MasterUserType.UserType)
	}

	var account entity.MasterAccount
	_ = db.Model(&entity.MasterAccount{}).Association("Role").Find(&account)

	fmt.Println(account)
}
