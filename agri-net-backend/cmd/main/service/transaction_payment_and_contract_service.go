package service

import (
	"context"
	"strconv"
	"time"

	tm "github.com/buger/goterm"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/contract"
	"github.com/samuael/agri-net/agri-net-backend/pkg/payment"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
)

// TransactionPaymentAndContractRoutine
type TransactionPaymentAndContractRoutine struct {
	PaymentService  payment.IPaymentService
	UserService     user.IUserService
	ContractService contract.IContractService
}

// NewTPACRoutine
func NewTPACRoutine(
	paymentservice payment.IPaymentService,
	userService user.IUserService,
	contractService contract.IContractService,
) *TransactionPaymentAndContractRoutine {
	return &TransactionPaymentAndContractRoutine{
		PaymentService:  paymentservice,
		UserService:     userService,
		ContractService: contractService,
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
				for x := range payments {
					payment := payments[x]
					deletePayment := false
					// check ther expiry time
					if !(payment.KebdCompleted) {
						binvoice, er := run.PaymentService.GetInvoiceByID(context.Background(), payment.BuyerInvoiceID)
						if (binvoice != nil) && (binvoice.Status == string(model.PROCESSED)) {
							payment.GuaranteeCompleted = true
						} else if er != nil || binvoice == nil || binvoice.Status == string(model.CANCELED) ||
							binvoice.Status == string(model.FAILED) ||
							binvoice.Status == string(model.EXPIRED) ||
							binvoice.Status == string(model.NONE) {
							if payment.GuaranteeCompleted {
								ctx := context.Background()
								ctx = context.WithValue(ctx, "user_id", uint64(payment.SellerID))
								user, _, _, ers := run.UserService.GetUserByEmailOrID(ctx)
								if ers != nil {
									continue
								}
								// refund the money
								vreq := &model.HellocashTransfferRequest{
									To:       user.Phone,
									Currency: "ETB",
									Amount:   int(payment.GuaranteeAmount),
								}
								_, er := run.PaymentService.ValidateMoneyTransffer(context.Background(), vreq)
								if er != nil {
									println("The Error is ", er.Error())
									continue
								}
								_, er = run.PaymentService.MoneyTransffer(context.Background(), vreq)
								if er != nil {
									println("The Error is ", er.Error())
									continue
								}
							}
							deletePayment = true
						}
						if deletePayment {
							run.PaymentService.DeletePaymentByTransactionID(context.Background(), payment.TransactionID)
						}
					}
					if !(payment.GuaranteeCompleted) {
						sinvoice, ers := run.PaymentService.GetInvoiceByID(context.Background(), payment.SellerInvoiceID)
						if ers != nil {
							continue
						}
						if (sinvoice != nil) && (sinvoice.Status == string(model.PROCESSED)) {
							payment.GuaranteeCompleted = true
						} else if er != nil || sinvoice == nil || sinvoice.Status == string(model.CANCELED) ||
							sinvoice.Status == string(model.FAILED) ||
							sinvoice.Status == string(model.EXPIRED) ||
							sinvoice.Status == string(model.NONE) {

							if payment.KebdCompleted {
								ctx := context.Background()
								ctx = context.WithValue(ctx, "user_id", uint64(payment.BuyerID))
								user, _, _, ers := run.UserService.GetUserByEmailOrID(ctx)
								if ers != nil {
									continue
								}
								vreq := &model.HellocashTransfferRequest{
									To:       user.Phone,
									Currency: "ETB",
									Amount:   int(payment.KebdAmount),
								}
								_, er := run.PaymentService.ValidateMoneyTransffer(context.Background(), vreq)
								if er != nil {
									continue
								}
								_, er = run.PaymentService.MoneyTransffer(context.Background(), vreq)
								if er != nil {
									continue
								}
							}
							deletePayment = true
						}
						if deletePayment {
							run.PaymentService.DeletePaymentByTransactionID(context.Background(), payment.TransactionID)
						}
					}
					if payment.KebdCompleted && payment.GuaranteeCompleted {
						secret := helper.GenerateRandomString(5, helper.NUMBERS)
						run.ContractService.CreateContract(context.Background(), int(payment.TransactionID), secret)
					} else if payment.CreatedAt < uint64(time.Now().Unix()-(60*90)) {
						if payment.GuaranteeCompleted {
							ctx := context.Background()
							ctx = context.WithValue(ctx, "user_id", uint64(payment.SellerID))
							user, _, _, ers := run.UserService.GetUserByEmailOrID(ctx)
							if ers != nil {
								continue
							}
							vreq := &model.HellocashTransfferRequest{
								To:       user.Phone,
								Currency: "ETB",
								Amount:   int(payment.GuaranteeAmount),
							}
							_, er := run.PaymentService.ValidateMoneyTransffer(context.Background(), vreq)
							if er != nil {
								continue
							}
							_, er = run.PaymentService.MoneyTransffer(context.Background(), vreq)
							if er != nil {
								continue
							}
						}
						// -----------------------------------------------------------------
						if payment.KebdCompleted {
							ctx := context.Background()
							ctx = context.WithValue(ctx, "user_id", uint64(payment.BuyerID))
							user, _, _, ers := run.UserService.GetUserByEmailOrID(ctx)
							if ers != nil {
								continue
							}
							vreq := &model.HellocashTransfferRequest{
								To:       user.Phone,
								Currency: "ETB",
								Amount:   int(payment.KebdAmount),
							}
							_, er := run.PaymentService.ValidateMoneyTransffer(context.Background(), vreq)
							if er != nil {
								continue
							}
							_, er = run.PaymentService.MoneyTransffer(context.Background(), vreq)
							if er != nil {
								continue
							}
						}
						// -------------------------------------------------
						run.PaymentService.DeletePaymentByTransactionID(context.Background(), payment.TransactionID)
					}
				}
			}

		}
	}
}
