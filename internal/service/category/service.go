package category

import (
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
