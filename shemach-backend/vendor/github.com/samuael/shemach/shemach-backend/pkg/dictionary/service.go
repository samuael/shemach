package dictionary

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IDictionaryService interface {
	NewDictionary(ctx context.Context, dict *model.Dictionary) error
	Translate(ctx context.Context, dict *model.Dictionary) error
	DeleteTranslation(ctx context.Context, dict *model.Dictionary) (int, error)
	UpdateTranslation(ctx context.Context, dict *model.Dictionary) error
	GetDictionaries(ctx context.Context, offset, limit uint) ([]*model.Dictionary, error)
}

type DictionaryService struct {
	Repo IDictionaryRepo
}

func NewDictionaryService(repo IDictionaryRepo) IDictionaryService {
	return &DictionaryService{
		Repo: repo,
	}
}

func (service *DictionaryService) NewDictionary(ctx context.Context, dict *model.Dictionary) error {
	return service.Repo.NewDictionary(ctx, dict)
}

func (service *DictionaryService) Translate(ctx context.Context, dict *model.Dictionary) error {
	return service.Repo.Translate(ctx, dict)
}
func (service *DictionaryService) DeleteTranslation(ctx context.Context, dict *model.Dictionary) (int, error) {
	return service.Repo.DeleteTranslation(ctx, dict)
}
func (service *DictionaryService) UpdateTranslation(ctx context.Context, dict *model.Dictionary) error {
	return service.Repo.UpdateTranslation(ctx, dict)
}

func (service *DictionaryService) GetDictionaries(ctx context.Context, offset, limit uint) ([]*model.Dictionary, error) {
	return service.Repo.GetDictionaries(ctx, offset, limit)
}
