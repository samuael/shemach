package crop

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ICropRepo interface {
	CreateCrop(ctx context.Context, crop *model.Crop) (int, error)
}
