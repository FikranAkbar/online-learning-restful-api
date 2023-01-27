package payment_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
	"online-learning-restful-api/config"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/payment"
	"online-learning-restful-api/repository/payment_repository"
)

type PaymentServiceImpl struct {
	payment_repository.PaymentRepository
	*validator.Validate
	*gorm.DB
}

func NewPaymentServiceImpl(paymentRepository payment_repository.PaymentRepository, validate *validator.Validate, DB *gorm.DB) *PaymentServiceImpl {
	return &PaymentServiceImpl{PaymentRepository: paymentRepository, Validate: validate, DB: DB}
}

func (service *PaymentServiceImpl) CreateNewCourseOrder(ctx context.Context, courseIds []uint) payment.CourseOrderResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	conf := config.LoadConfigFromEnv()

	var s = snap.Client{
		ServerKey: conf.MidtransServerKey,
		Env:       midtrans.Sandbox,
	}
	paymentHistory, courses, user, err := service.PaymentRepository.CreateNewCourseOrder(ctx, tx, courseIds)
	helper.PanicIfError(err)

	snapReq := helper.GenerateSnapReq(paymentHistory, user, courses)

	_, err = s.CreateTransaction(snapReq)
	helper.PanicIfError(err)

	var itemDetailsResponse []payment.ItemDetailResponse
	for _, course := range courses {
		itemDetailsResponse = append(itemDetailsResponse, payment.ItemDetailResponse{
			Id:    course.Id,
			Name:  course.Name,
			Price: course.Price,
		})
	}

	courseOrderResponse := payment.CourseOrderResponse{
		OrderId:     paymentHistory.OrderId,
		ItemDetails: itemDetailsResponse,
	}

	return courseOrderResponse
}
