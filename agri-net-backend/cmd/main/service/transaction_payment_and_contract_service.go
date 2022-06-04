package service

import (
	"context"
	"strconv"
	"time"

	tm "github.com/buger/goterm"

	"github.com/samuael/agri-net/agri-net-backend/pkg/payment"
)

// TransactionPaymentAndContractRoutine
type TransactionPaymentAndContractRoutine struct {
	PaymentService payment.IPaymentService
}

// NewTPACRoutine
func NewTPACRoutine(
	paymentservice payment.IPaymentService) *TransactionPaymentAndContractRoutine {
	return &TransactionPaymentAndContractRoutine{
		PaymentService: paymentservice,
	}
}

func (run *TransactionPaymentAndContractRoutine) Run() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			{
				payments, er := run.PaymentService.GetPendingPayment(context.Background())
				tm.Println(tm.Color("Pending Payments: "+strconv.Itoa(len(payments)), tm.GREEN))
				if er != nil {
					println(er.Error())
				}
			}

		}
	}
}
