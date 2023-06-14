package category

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-category/internal/entity"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
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

func (s *Service) GetCategory(categoryId uuid.UUID, userId uuid.UUID) (entity.Category, error) {
	category, ownerId, err := s.repo.GetCategory(categoryId)
	if err != nil {
		return entity.Category{}, err
	}

	if userId != ownerId {
		return entity.Category{}, fail.GrpcAccessDenied
	}

	return category, nil
}
func (s *Service) UpdateCategory(category entity.Category, userId uuid.UUID) error {
	return s.repo.UpdateCategory(category, userId)
}

func (s *Service) DeleteCategory(categoryId, userId uuid.UUID) error {
	return s.repo.DeleteCategory(categoryId, userId)
}
