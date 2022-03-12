package author

import (
	"context"
	"github.com/Masterminds/squirrel"
)

type Repository interface {
	Create(ctx context.Context, author *Author) error
	FindAll(ctx context.Context) ([]Author, error)
	FindOne(ctx context.Context, id string) (Author, error)
	Update(ctx context.Context, user Author) error
	Delete(ctx context.Context, id string) error
}

type SortOptions interface {
	EnrichQuery(q squirrel.SelectBuilder) squirrel.SelectBuilder
}
