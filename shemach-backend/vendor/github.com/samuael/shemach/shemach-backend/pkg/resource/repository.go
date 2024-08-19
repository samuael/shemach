package resource

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IResourceRepo interface {
	SaveImagesResources(ctx context.Context, resources []*model.PostImg) error
	GetImageByID(ctx context.Context, imgid uint64) (*model.PostImg, error)
}
