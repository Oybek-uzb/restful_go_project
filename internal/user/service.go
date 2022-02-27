package user

import (
	"context"
	"restful_go_project/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, e error) {
	// TODO later
	return
}
