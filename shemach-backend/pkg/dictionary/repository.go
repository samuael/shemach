package dictionary

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IDictionaryRepo interface {
	NewDictionary(ctx context.Context, dict *model.Dictionary) error
	Translate(ctx context.Context, dict *model.Dictionary) error
	DeleteTranslation(ctx context.Context, dict *model.Dictionary) (int, error)
	UpdateTranslation(ctx context.Context, dict *model.Dictionary) error
	GetDictionaries(ctx context.Context, offset, limit uint) ([]*model.Dictionary, error)
}
