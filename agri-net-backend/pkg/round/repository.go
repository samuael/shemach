package round

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IRoundRepository interface {
	GetRoundByRoundNumberAndCategoryID(ctx context.Context) (*model.Round, int, error)
	CreateRound(ctx context.Context) (*model.Round, int, error)
	DeleteRoundByID(ctx context.Context) error
	GetRoundByID(ctx context.Context) (*model.Round, error)
	UpdateRound(ctx context.Context) (int, error)
	GetRoundsOfCategory(context.Context) ([]*model.Round, int, error)
	CheckTheExistanceAndActivenessOfRound(context.Context) (int, int)
	CheckTheExistanceOfCategory(ctx context.Context) (int, int)
}
