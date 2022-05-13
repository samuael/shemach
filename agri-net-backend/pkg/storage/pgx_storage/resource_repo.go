package pgx_storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/resource"
)

type ResourceRepo struct {
	DB *pgxpool.Pool
}

func NewResourceRepo(conn *pgxpool.Pool) resource.IResourceService {
	return &ResourceRepo{}
}
