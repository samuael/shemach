package contract

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IContractService interface {
	CreateContract(ctx context.Context, transactionid int, secret string) (*model.Contract, error)
	GetContractByTransactionID(ctx context.Context, txid uint64) (*model.Contract, error)
}

type ContractService struct {
	Repo IContractRepo
}

func NewContractService(repo IContractRepo) IContractService {
	return &ContractService{
		Repo: repo,
	}
}
func (ser *ContractService) CreateContract(ctx context.Context, transactionid int, secret string) (*model.Contract, error) {
	return ser.Repo.CreateContract(ctx, transactionid, secret)
}
func (ser *ContractService) GetContractByTransactionID(ctx context.Context, txid uint64) (*model.Contract, error) {
	return ser.Repo.GetContractByTransactionID(ctx, txid)
}
