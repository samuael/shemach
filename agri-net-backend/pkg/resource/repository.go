package resource

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IResourceRepo interface {
	SaveImagesResources(ctx context.Context, resources []*model.PostImg) error
}
