package cart_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"strconv"
	"strings"
)

type CartRepositoryImpl struct {
}

func NewCartRepositoryImpl() *CartRepositoryImpl {
	return &CartRepositoryImpl{}
}

func (repository *CartRepositoryImpl) BuyCartItems(ctx context.Context, db *gorm.DB, courseIds []uint) (domain.PaymentHistory, []domain.Course, domain.User, error) {
	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	var userEntity entity.MasterUser
	err := db.WithContext(ctx).
		Where("id = ?", userTokenInfo.UserId).
		First(&userEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("User with id %v not found", userTokenInfo.UserId)
		return domain.PaymentHistory{}, nil, domain.User{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.PaymentHistory{}, nil, domain.User{}, err
	}

	var courseEntities []entity.MasterCourse
	var courses []domain.Course
	var totalCoursesPrice uint
	for _, courseId := range courseIds {
		var courseEntity entity.MasterCourse
		err = db.WithContext(ctx).
			Where("id = ?", courseId).
			First(&courseEntity).Error
		if err != nil && exception.CheckErrorContains(err, exception.NotFound) || !courseEntity.IsPublished {
			logError := fmt.Sprintf("Course with id %v not found", courseId)
			return domain.PaymentHistory{}, nil, domain.User{}, exception.GenerateHTTPError(exception.NotFound, logError)
		} else if err != nil {
			return domain.PaymentHistory{}, nil, domain.User{}, err
		}

		courseEntities = append(courseEntities, courseEntity)

		course := domain.Course{
			Id:    courseEntity.ID,
			Name:  courseEntity.Name,
			Price: courseEntity.Price,
		}
		courses = append(courses, course)

		totalCoursesPrice = totalCoursesPrice + courseEntity.Price
	}

	courseIdsInString := strings.Trim(strings.Join(strings.Split(fmt.Sprint(courseIds), " "), ","), "[]")

	paymentHistoryEntity := entity.TrxUserPaymentHistory{
		OrderId:         helper.RandStringBytes(10),
		UserId:          userTokenInfo.UserId,
		CourseIds:       courseIdsInString,
		DayId:           1,
		PaymentStatusId: 3,
		PaymentMethodId: 1,
		Price:           totalCoursesPrice,
		TotalPrice:      totalCoursesPrice,
		IsExpired:       false,
	}
	err = db.WithContext(ctx).Create(&paymentHistoryEntity).Error
	if err != nil {
		return domain.PaymentHistory{}, nil, domain.User{}, err
	}

	for _, courseId := range courseIds {
		var userCart entity.TrxUserCart
		rowAffected := db.WithContext(ctx).
			Model(&userCart).
			Where("user_id = ?", userTokenInfo.UserId).
			Where("course_id = ?", courseId).
			Delete(&userCart).RowsAffected
		if rowAffected <= 0 {
			return domain.PaymentHistory{}, nil, domain.User{},
				exception.GenerateHTTPError(exception.BadRequest, "The selected course is not found in user cart")
		}
	}

	return domain.PaymentHistory{
			OrderId:         paymentHistoryEntity.OrderId,
			UserId:          userTokenInfo.UserId,
			CourseIds:       courseIdsInString,
			DayId:           paymentHistoryEntity.DayId,
			PaymentStatusId: paymentHistoryEntity.PaymentStatusId,
			TotalPrice:      paymentHistoryEntity.TotalPrice,
			PaymentMethodId: paymentHistoryEntity.PaymentMethodId,
			PaymentUrl:      paymentHistoryEntity.PaymentUrl,
			IsExpired:       paymentHistoryEntity.IsExpired,
		}, courses, domain.User{
			Id:        userEntity.ID,
			Name:      userEntity.Name,
			Email:     userTokenInfo.UserEmail,
			BirthDate: userEntity.BirthDate.Time,
			Gender:    userEntity.Gender,
			Phone:     userEntity.Phone.String,
		}, nil
}

func (repository *CartRepositoryImpl) SavePaymentRedirectionURL(ctx context.Context, db *gorm.DB, paymentHistory domain.PaymentHistory) (domain.PaymentHistory, error) {
	var paymentHistoryEntity entity.TrxUserPaymentHistory
	err := db.WithContext(ctx).
		Model(&paymentHistoryEntity).
		Where("order_id = ?", paymentHistory.OrderId).
		Updates(map[string]any{
			"payment_url": paymentHistory.PaymentUrl,
		}).First(&paymentHistoryEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("payment history with id %v not found", paymentHistory.OrderId)
		return domain.PaymentHistory{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.PaymentHistory{}, err
	}

	return domain.PaymentHistory{
		OrderId:         paymentHistoryEntity.OrderId,
		UserId:          paymentHistoryEntity.UserId,
		CourseIds:       paymentHistoryEntity.CourseIds,
		DayId:           paymentHistoryEntity.DayId,
		PaymentStatusId: paymentHistoryEntity.PaymentStatusId,
		TotalPrice:      paymentHistoryEntity.TotalPrice,
		PaymentMethodId: paymentHistoryEntity.PaymentMethodId,
		PaymentUrl:      paymentHistoryEntity.PaymentUrl,
		IsExpired:       false,
	}, nil
}

func (repository *CartRepositoryImpl) AddNewCourseToUser(ctx context.Context, db *gorm.DB, paymentNotification domain.MidtransPaymentNotification) error {
	var paymentHistoryEntity entity.TrxUserPaymentHistory
	err := db.WithContext(ctx).
		Model(&paymentHistoryEntity).
		Where("order_id = ?", paymentNotification.OrderId).
		Updates(map[string]any{"payment_status_id": 1}).
		First(&paymentHistoryEntity).Error
	if err != nil {
		return err
	}

	for _, courseIdString := range strings.Split(paymentHistoryEntity.CourseIds, ",") {
		courseId, err := strconv.Atoi(courseIdString)
		if err != nil {
			break
		}

		userCourse := domain.UserCourse{
			UserId:   paymentHistoryEntity.UserId,
			CourseId: uint(courseId),
		}

		userCourseEntity := entity.TrxUserCourse{
			UserId:               userCourse.UserId,
			CourseId:             userCourse.CourseId,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 1,
		}
		err = db.WithContext(ctx).
			Create(&userCourseEntity).Error
		if err != nil {
			break
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (repository *CartRepositoryImpl) ChangeUserPaymentHistoryToCancelled(ctx context.Context, db *gorm.DB, paymentNotification domain.MidtransPaymentNotification) error {
	var paymentHistoryEntity entity.TrxUserPaymentHistory
	err := db.WithContext(ctx).
		Model(&paymentHistoryEntity).
		Where("order_id = ?", paymentNotification.OrderId).
		Updates(map[string]any{"payment_status_id": 2}).
		First(&paymentHistoryEntity).Error
	if err != nil {
		return err
	}

	return nil
}
