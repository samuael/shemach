package resource

import (
	"context"

	"github.com/samuael/shemach/backend/pkg/constants/types"
)

type IResourceRepo interface {
	SaveImagesResources(ctx context.Context, resources []*types.PostImg) error
	GetImageDetailByID(ctx context.Context, imgid uint64) (*types.PostImg, error)
	GetImagePath(ctx context.Context, imgid uint64) (string, error)

	CheckPropertyImageUploadEligibility(ctx context.Context, propertyDetailID, createdBy uint64, isCoverPicture bool, maxImages uint8) (uint64, error)
	SavePropertyImage(ctx context.Context, path string, propertyID, createdBy uint64, isCoverPicture bool, role uint8) (uint64, error)
}
