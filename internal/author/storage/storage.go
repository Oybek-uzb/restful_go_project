package storage

import (
	"context"
	"restful_go_project/internal/author/model"
)

type Repository interface {
	Create(ctx context.Context, author *model.Author) error
	FindAll(ctx context.Context, sortOptions SortOptions) ([]model.Author, error)
	FindOne(ctx context.Context, id string) (model.Author, error)
	Update(ctx context.Context, user model.Author) error
	Delete(ctx context.Context, id string) error
}
