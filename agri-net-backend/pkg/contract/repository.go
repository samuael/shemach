package contract

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IContractRepo interface {
	CreateContract(ctx context.Context, transactionid int, secret string) (*model.Contract, error)
	GetContractByTransactionID(ctx context.Context, txid uint64) (*model.Contract, error)
}
