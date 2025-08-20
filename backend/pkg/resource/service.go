package resource

import (
	"context"

	"github.com/samuael/shemach/backend/pkg/constants/types"
)

// IResourceService interface representing the main crop type
type IResourceService interface {
	SaveImagesResources(ctx context.Context, resources []*types.PostImg) error
	GetImageDetailByID(ctx context.Context, imgid uint64) (*types.PostImg, error)
	GetImagePath(ctx context.Context, imgid uint64) (string, error)

	// new
	CheckPropertyImageUploadEligibility(ctx context.Context, propertyDetailID, createdBy uint64, isCoverPicture bool, maxImages uint8) (uint64, error)

	// SavePropertyImage saves an image information into the property detail
	// -- 1 if the property not found
	// -- 2 owner of the property mismatch mismatch
	// -- 3 insertion error
	// -- 4 cover picture update error
	// -- 5 pictures update error
	// -- bigint successful insertion.
	SavePropertyImage(ctx context.Context, path string, propertyID, createdBy uint64, isCoverPicture bool, role uint8) (uint64, error)
}

// ResourceService ...
type ResourceService struct {
	Repo IResourceRepo
}

// NewResourceService
func NewResourceService(repo IResourceRepo) IResourceService {
	return &ResourceService{
		Repo: repo,
	}
}

func (service *ResourceService) SaveImagesResources(ctx context.Context, resources []*types.PostImg) error {
	return service.Repo.SaveImagesResources(ctx, resources)
}

func (service *ResourceService) GetImageDetailByID(ctx context.Context, imgid uint64) (*types.PostImg, error) {
	return service.Repo.GetImageDetailByID(ctx, imgid)
}

func (service *ResourceService) CheckPropertyImageUploadEligibility(ctx context.Context, propertyDetailID, createdBy uint64, isCoverPicture bool, maxImages uint8) (uint64, error) {
	return service.Repo.CheckPropertyImageUploadEligibility(ctx, propertyDetailID, createdBy, isCoverPicture, maxImages)
}

func (service *ResourceService) SavePropertyImage(ctx context.Context, path string, propertyID, createdBy uint64, isCoverPicture bool, role uint8) (uint64, error) {
	return service.Repo.SavePropertyImage(ctx, path, propertyID, createdBy, isCoverPicture, role)
}

func (service *ResourceService) GetImagePath(ctx context.Context, imgid uint64) (string, error) {
	return service.Repo.GetImagePath(ctx, imgid)
}
