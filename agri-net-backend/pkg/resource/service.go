package resource

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

// IResourceService interface representing the main crop type
type IResourceService interface {
	SaveImagesResources(ctx context.Context, resources []*model.PostImg) error
}

// ResourceService ...
type ResourceService struct {
	Repo IResourceService
}

// NewResourceService
func NewResourceService(repo IResourceRepo) IResourceService {
	return &ResourceService{
		Repo: repo,
	}
}

func (service *ResourceService) SaveImagesResources(ctx context.Context, resources []*model.PostImg) error {
	return service.Repo.SaveImagesResources(ctx, resources)
}
