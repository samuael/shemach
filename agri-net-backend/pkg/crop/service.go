package crop

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

// ICropService interface representing the main crop type
type ICropService interface {
	CreateCrop(ctx context.Context, crop *model.Crop) (int, error)
	GetPostByID(ctx context.Context, postid uint64) (*model.Crop, error)
	SaveNewPostImages(ctx context.Context, postif uint64, images []int) error
	GetPosts(ctx context.Context, offset, limit uint) ([]*model.Crop, error)
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

// GetPostByID(ctx context.Context, postid uint64) (*model.Crop, error)
func (service *CropService) GetPostByID(ctx context.Context, postid uint64) (*model.Crop, error) {
	return service.Repo.GetPostByID(ctx, postid)
}
func (service *CropService) SaveNewPostImages(ctx context.Context, postid uint64, images []int) error {
	return service.Repo.SaveNewPostImages(ctx, postid, images)
}

func (service *CropService) GetPosts(ctx context.Context, offset, limit uint) ([]*model.Crop, error) {
	return service.Repo.GetPosts(ctx, offset, limit)
}
