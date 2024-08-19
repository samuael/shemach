package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/contract"
)

// ContractRepo
type ContractRepo struct {
	DB *pgxpool.Pool
}

// NewContractRepo
func NewContractRepo(conn *pgxpool.Pool) contract.IContractRepo {
	return &ContractRepo{
		DB: conn,
	}
}

// CreateContract
func (repo *ContractRepo) CreateContract(ctx context.Context, transactionid int, secret string) (*model.Contract, error) {
	state := 0
	er := repo.DB.QueryRow(ctx, "select * from createContract($1,$2)", transactionid, secret).Scan(&state)
	if er != nil {
		return nil, er
	} else if state <= 0 {
		return nil, errors.New("could not create contract")
	}
	return &model.Contract{
		ID:            uint64(state),
		TransactionID: uint64(transactionid),
		SecretString:  secret,
		State:         18,
	}, er
}

// GetContractByTransactionID
func (repo *ContractRepo) GetContractByTransactionID(ctx context.Context, txid uint64) (*model.Contract, error) {
	contract := &model.Contract{}
	er := repo.DB.QueryRow(ctx, `select contract_id ,transaction_id ,secret_string ,state from contract where transaction_id=$1`, txid).
		Scan(contract.ID, contract.TransactionID, contract.SecretString, contract.State)
	if er != nil {
		return nil, er
	}
	return contract, nil
}
