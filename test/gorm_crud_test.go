package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	"online-learning-restful-api/app/database"
	entity2 "online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/helper"
	"testing"
	"time"
)

func TestFirstData(t *testing.T) {
	db := database.NewDB()
	var account *entity2.MasterAccount
	db.Preload("MasterUserType").First(&account)
	fmt.Print(helper.StringModel(*account))
}

func TestFindDataAndRelationshipData(t *testing.T) {
	db := database.NewDB()
	var accounts []entity2.MasterAccount
	db.Preload(clause.Associations).Find(&accounts)

	for _, val := range accounts {
		fmt.Printf("Name %v Role %v\n", val.Email, val.MasterUserType.Name)
	}
}

func TestFindDataManyToMany(t *testing.T) {
	db := database.NewDB()
	var users []entity2.MasterUser
	db.Model(&entity2.MasterUser{}).Preload(clause.Associations).Find(&users)

	for _, user := range users {
		fmt.Printf("My name is %v and my course is", user.Name)

		for _, course := range user.Courses {
			fmt.Printf(" %v", course.Name)
		}

		fmt.Println()
	}
}

func TestTimeVariable(t *testing.T) {
	result, _ := time.Parse("03:04", "09:11")
	fmt.Println(result)
}

func TestBcryptPassword(t *testing.T) {
	hashedByte, _ := bcrypt.GenerateFromPassword([]byte("Password"), bcrypt.DefaultCost)
	fmt.Println(string(hashedByte))
}
