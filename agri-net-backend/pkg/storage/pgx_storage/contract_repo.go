package pgx_storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/contract"
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
