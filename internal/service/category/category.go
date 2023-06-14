package category

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-category/internal/entity"
)

func (s *Service) GetUserCategories(userId uuid.UUID) []entity.Category {
	return s.repo.GetUserCategories(userId)
}

func (s *Service) GetCategoriesMap(categoryIds []uuid.UUID, userId uuid.UUID) map[uuid.UUID]entity.Category {
	return s.repo.GetCategoriesMap(categoryIds, userId)
}

func (s *Service) AddCategory(category entity.Category, userId uuid.UUID) (uuid.UUID, error) {
	return s.repo.AddCategory(category, userId)
}
func (s *Service) UpdateCategory(category entity.Category, userId uuid.UUID) error {
	return s.repo.UpdateCategory(category, userId)
}

func (s *Service) DeleteCategory(categoryId, userId uuid.UUID) error {
	return s.repo.DeleteCategory(categoryId, userId)
}
