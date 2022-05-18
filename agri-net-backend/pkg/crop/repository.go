package crop

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ICropRepo interface {
	CreateCrop(ctx context.Context, crop *model.Crop) (int, error)
	GetPostByID(ctx context.Context, postid uint64) (*model.Crop, error)
	SaveNewPostImages(ctx context.Context, postid uint64, images []int) error
}
