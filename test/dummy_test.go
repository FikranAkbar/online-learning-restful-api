package test

import (
	"context"
	"fmt"
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/helper"
	"testing"
)

func TestName(t *testing.T) {
	db := database.NewDB()
	var userCourseEntity entity.TrxUserCourse
	err := db.WithContext(context.Background()).
		Where("user_id = ?", 1).
		Where("course_id = ?", 6).
		Preload("User").
		Preload("Course").
		First(&userCourseEntity).Error

	helper.PanicIfError(err)

	fmt.Println(userCourseEntity.User.Name)
	fmt.Println(userCourseEntity.Course.Name)
}
