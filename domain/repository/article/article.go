package article

import (
	"context"
	"m1-article-service/domain/entity"
)

type Article interface {
	Create(context.Context, *entity.Article) (int64, error)
	Update(context.Context, *entity.Article) error
	Delete(context.Context, int64) error
	Detail(context.Context, int64) (*entity.Article, error)
	List(context.Context, uint16) ([]*entity.Article, error)
}
