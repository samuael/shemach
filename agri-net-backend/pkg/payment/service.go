package payment

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IPaymentService interface {
	// CheckTheExistanceOfPaymentInstance uses "payin_input" of type *model.PayinInput to return a status information as int.
	// -4 internal database query error
	// -3 the round has not the specified student
	// -2 round doesnt exist
	// -1 student doesnt exist
	//  0 payment information doesnt exist
	//  1 payment information with similar data already exist'
	CheckTheExistanceOfPaymentInstance(ctx context.Context) int
	// GetPaymentUsingPayinInput uses "payin_input" of type *model.PayinInput to return *model.PayIn, status code (int), error
	GetPaymentUsingPayinInput(ctx context.Context) (*model.PayIn, int, error)
	// CreatePayinInstance uses "*payin_input" of type model.PayinInput and returns *model.PayIn, "status code" int, error
	CreatePayinInstance(ctx context.Context) (*model.PayIn, int, error)
	// GetRemainingPaymentOfStudentForRound uses "student_id" of type string
	GetRemainingPaymentOfStudentForRound(ctx context.Context) (float64, int, error)
	// GetPaymentsOfAStudent uses "student_id" of type uint64 to return a list of payment transactions the student had made so far.
	GetPaymentsOfAStudent(ctx context.Context) ([]*model.PayIn, int, error)
}

type PaymentService struct {
	Repo IPaymentRepo
}

func NewPaymentService(repo IPaymentRepo) IPaymentService {
	return &PaymentService{
		Repo: repo,
	}
}

func (payser *PaymentService) CheckTheExistanceOfPaymentInstance(ctx context.Context) int {
	return payser.Repo.CheckTheExistanceOfPaymentInstance(ctx)
}

func (payser *PaymentService) GetPaymentUsingPayinInput(ctx context.Context) (*model.PayIn, int, error) {
	return payser.Repo.GetPaymentUsingPayinInput(ctx)
}

func (payser *PaymentService) CreatePayinInstance(ctx context.Context) (*model.PayIn, int, error) {
	return payser.Repo.CreatePayinInstance(ctx)
}

func (payser *PaymentService) GetRemainingPaymentOfStudentForRound(ctx context.Context) (float64, int, error) {
	return payser.Repo.GetRemainingPaymentOfStudentForRound(ctx)
}
func (payser *PaymentService) GetPaymentsOfAStudent(ctx context.Context) ([]*model.PayIn, int, error) {
	return payser.Repo.GetPaymentsOfAStudent(ctx)
}
