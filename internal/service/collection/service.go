package collection

import (
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
