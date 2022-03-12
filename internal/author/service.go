package author

import (
	"context"
	"fmt"
	"restful_go_project/pkg/api/sort"
	"restful_go_project/pkg/logging"
)

type Service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetAll(ctx context.Context, sortOptions sort.Options) ([]Author, error) {
	all, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all authors due to error: %v", err)
	}

	return all, nil
}
