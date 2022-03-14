package service

import (
	"context"
	"fmt"
	"restful_go_project/internal/author/model"
	"restful_go_project/internal/author/storage"
	model2 "restful_go_project/internal/author/storage/model"
	"restful_go_project/pkg/api/filter"
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

func (s *Service) GetAll(ctx context.Context, filterOptions filter.Options, sortOptions sort.Options) ([]model.Author, error) {

	options := model2.SortOptions{
		Field: sortOptions.Field,
		Order: sortOptions.Order,
	}
	all, err := s.repository.FindAll(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to get all authors due to error: %v", err)
	}

	return all, nil
}
