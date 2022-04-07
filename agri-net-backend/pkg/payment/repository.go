package payment

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IPaymentRepo interface {
	CheckTheExistanceOfPaymentInstance(ctx context.Context) int
	GetPaymentUsingPayinInput(ctx context.Context) (*model.PayIn, int, error)
	CreatePayinInstance(ctx context.Context) (*model.PayIn, int, error)
	GetRemainingPaymentOfStudentForRound(ctx context.Context) (float64, int, error)
	GetPaymentsOfAStudent(ctx context.Context) ([]*model.PayIn, int, error)
}
