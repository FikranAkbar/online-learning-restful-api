package cart_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
	"online-learning-restful-api/config"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"online-learning-restful-api/model/web/payment"
	"online-learning-restful-api/repository/cart_repository"
)

type CartServiceImpl struct {
	cart_repository.CartRepository
	*validator.Validate
	*gorm.DB
}

func NewCartServiceImpl(paymentRepository cart_repository.CartRepository, validate *validator.Validate, DB *gorm.DB) *CartServiceImpl {
	return &CartServiceImpl{CartRepository: paymentRepository, Validate: validate, DB: DB}
}

func (service *CartServiceImpl) BuyCartItems(ctx context.Context, courseIds []uint) payment.CourseOrderResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	conf := config.LoadConfigFromEnv()

	var s snap.Client
	s.New(conf.MidtransServerKey, midtrans.Sandbox)

	paymentHistory, courses, user, err := service.CartRepository.BuyCartItems(ctx, tx, courseIds)
	helper.PanicIfError(err)

	snapReq := helper.GenerateSnapReq(paymentHistory, user, courses)
	snapResp, _ := s.CreateTransaction(snapReq)
	paymentHistory.PaymentUrl = snapResp.RedirectURL

	_, err = service.CartRepository.SavePaymentRedirectionURL(ctx, tx, paymentHistory)
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

func (service *CartServiceImpl) HandleMidtransPaymentNotification(ctx context.Context, request payment.MidtransPaymentNotificationRequest) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if request.TransactionStatus == "settlement" {
		paymentNotification := domain.MidtransPaymentNotification{
			TransactionStatus: request.TransactionStatus,
			OrderId:           request.OrderId,
			PaymentType:       request.PaymentType,
		}
		err := service.CartRepository.AddNewCourseToUser(ctx, tx, paymentNotification)
		helper.PanicIfError(err)
	} else if request.TransactionStatus == "cancel" ||
		request.TransactionStatus == "deny" ||
		request.TransactionStatus == "expire" {
		paymentNotification := domain.MidtransPaymentNotification{
			TransactionStatus: request.TransactionStatus,
			OrderId:           request.OrderId,
			PaymentType:       request.PaymentType,
		}
		err := service.CartRepository.ChangeUserPaymentHistoryToCancelled(ctx, tx, paymentNotification)
		helper.PanicIfError(err)
	}
}
