package helper

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"online-learning-restful-api/model/domain"
	"strconv"
)

func GenerateSnapReq(paymentHistory domain.PaymentHistory, user domain.User, courses []domain.Course) *snap.Request {
	custDetail := midtrans.CustomerDetails{
		FName: user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	var items []midtrans.ItemDetails
	var grossAmount int64
	for _, course := range courses {
		items = append(items, midtrans.ItemDetails{
			ID:    strconv.Itoa(int(course.Id)),
			Name:  course.Name,
			Price: int64(course.Price),
			Qty:   1,
		})

		grossAmount = grossAmount + int64(course.Price)
	}

	return &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentHistory.OrderId,
			GrossAmt: grossAmount,
		},
		CustomerDetail: &custDetail,
		Items:          &items,
	}
}
