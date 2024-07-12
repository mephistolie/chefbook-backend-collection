package mq

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-collection/internal/service/dependencies/repository"
)

type Service struct {
	repo repository.Collection
}

func NewService(
	repo repository.Collection,
) *Service {
	return &Service{repo: repo}
}

func (s *Service) DeleteUser(userId uuid.UUID, messageId uuid.UUID) error {
	return s.repo.DeleteUser(userId, messageId)
}
