package crop

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

// ICropService interface representing the main crop type
type ICropService interface {
	CreateCrop(ctx context.Context, crop *model.Crop) (int, error)
}

// CropService ...
type CropService struct {
	Repo ICropRepo
}

func NewCropService(repo ICropRepo) ICropService {
	return &CropService{
		Repo: repo,
	}
}

// CreateCrop
func (service *CropService) CreateCrop(ctx context.Context, crop *model.Crop) (int, error) {
	return service.Repo.CreateCrop(ctx, crop)
}
