package service

import (
	"context"
	"fmt"
	"restful_go_project/internal/author/model"
	"restful_go_project/internal/author/storage"
	"restful_go_project/pkg/api/sort"
	"restful_go_project/pkg/logging"
)

type Service struct {
	repository storage.Repository
	logger     *logging.Logger
}

func NewService(repository storage.Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetAll(ctx context.Context, sortOptions sort.Options) ([]model.Author, error) {
	options := storage.SortOptions{
		Field: sortOptions.Field,
		Order: sortOptions.Order,
	}
	all, err := s.repository.FindAll(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to get all authors due to error: %v", err)
	}

	return all, nil
}
