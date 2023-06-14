package mq

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-category/internal/service/dependencies/repository"
)

type Service struct {
	repo repository.Category
}

func NewService(
	repo repository.Category,
) *Service {
	return &Service{repo: repo}
}

func (s *Service) DeleteUser(userId uuid.UUID, messageId uuid.UUID) error {
	return s.repo.DeleteUser(userId, messageId)
}
