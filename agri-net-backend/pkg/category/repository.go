package category

import (
	"context"

	"github.com/samuael/Project/RegistrationSystem/pkg/constants/model"
)

// ICategoryRepo
type ICategoryRepo interface {
	Create(context.Context) (*model.Category, error, int)
	GetCategoryByID(ctx context.Context) (*model.Category, error, int)
	GetCategories(ctx context.Context) ([]*model.Category, error)
	DeleteCategoryByID(ctx context.Context) error
	UpdateCategory(ctx context.Context) (error, int)
	GetCategoryByContent(ctx context.Context) *model.Category
	GetImageUrl(ctx context.Context) (string, error)
	ChangeCategoryImageUrl(ctx context.Context) error
	GetCategoryStudents(ctx context.Context) (*model.CategoryStudentCountReponse, int, error)
}
