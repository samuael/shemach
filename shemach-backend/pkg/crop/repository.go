package crop

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type ICropRepo interface {
	CreateCrop(ctx context.Context, crop *model.Crop) (int, error)
	GetPostByID(ctx context.Context, postid uint64) (*model.Crop, error)
	SaveNewPostImages(ctx context.Context, postid uint64, images []int) error
	GetPosts(ctx context.Context, offset, limit uint) ([]*model.Crop, error)
	GetAgentPosts(ctx context.Context, userid uint64, offset, limit uint) ([]*model.Crop, error)
	GetMerchantPosts(ctx context.Context, storeIDS []uint64, offset, limit uint) ([]*model.Crop, error)
}
