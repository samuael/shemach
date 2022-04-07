package category

import (
	"context"

	"github.com/samuael/Project/RegistrationSystem/pkg/constants/model"
)

type ICategoryService interface {
	// GetCategoryByContent uses 'category' in the context toget any category in the dataase with same title and short title.
	GetCategoryByContent(context.Context) *model.Category
	// Create ... 'category'
	Create(context.Context) (*model.Category, error, int)
	// GetCategoryByID ...  uses 'category_id' type uint64  to get a category
	GetCategoryByID(ctx context.Context) (*model.Category, error, int)
	// GetCategories returns all the categories in the database.
	GetCategories(ctx context.Context) ([]*model.Category, error)
	// DeleteCategoryByID uses 'category_id' uint64 to delete the category and returns error if error has happened
	// returns nil otherwise.
	DeleteCategoryByID(ctx context.Context) error
	// UpdateCategory uses 'category' of type *model.Category to update the category data
	UpdateCategory(ctx context.Context) (error, int)
	// GetImageUrl  uses "category_id" to search for a category instance and get the image url.
	GetImageUrl(ctx context.Context) (string, error)

	// ChangeCategoryImageUrl  uses "image_url" and "category_id"
	ChangeCategoryImageUrl(ctx context.Context) bool
	// GetCategoryStudents uses "category_id" type uint
	GetCategoryStudents(ctx context.Context) (*model.CategoryStudentCountReponse, int, error)
}

type CategoryService struct {
	Repo ICategoryRepo
}

func NewCategoryService(repo ICategoryRepo) ICategoryService {
	return &CategoryService{Repo: repo}
}

// GetCategoryByContent ...
func (catser *CategoryService) GetCategoryByContent(ctx context.Context) *model.Category {
	return catser.Repo.GetCategoryByContent(ctx)
}

// Create takes
func (catser *CategoryService) Create(ctx context.Context) (*model.Category, error, int) {
	return catser.Repo.Create(ctx)
}

// GetCategoryByID
func (catser *CategoryService) GetCategoryByID(ctx context.Context) (*model.Category, error, int) {
	return catser.Repo.GetCategoryByID(ctx)
}

// GetCategories ...
func (catser *CategoryService) GetCategories(ctx context.Context) ([]*model.Category, error) {
	return catser.Repo.GetCategories(ctx)
}

// DeleteCategoryByID
func (catser *CategoryService) DeleteCategoryByID(ctx context.Context) error {
	return catser.Repo.DeleteCategoryByID(ctx)
}

// UpdateCategory ... 'category'
func (catser *CategoryService) UpdateCategory(ctx context.Context) (error, int) {
	return catser.Repo.UpdateCategory(ctx)
}

// GetImageUrl
func (catser *CategoryService) GetImageUrl(ctx context.Context) (string, error) {
	return catser.Repo.GetImageUrl(ctx)
}

// ChangeCategoryImageUrl
func (catser *CategoryService) ChangeCategoryImageUrl(ctx context.Context) bool {
	return catser.Repo.ChangeCategoryImageUrl(ctx) == nil
}

func (catser *CategoryService) GetCategoryStudents(ctx context.Context) (*model.CategoryStudentCountReponse, int, error) {
	return catser.Repo.GetCategoryStudents(ctx)
}
