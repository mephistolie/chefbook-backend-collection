package repository

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-category/internal/entity"
)

type Category interface {
	DeleteUser(userId, messageId uuid.UUID) error

	GetUserCategories(userId uuid.UUID) []entity.Category
	GetCategoriesMap(categoryIds []uuid.UUID, userId uuid.UUID) map[uuid.UUID]entity.Category
	AddCategory(category entity.Category, userId uuid.UUID) (uuid.UUID, error)
	GetCategory(categoryId uuid.UUID) (entity.Category, uuid.UUID, error)
	UpdateCategory(category entity.Category, userId uuid.UUID) error
	DeleteCategory(categoryId, userId uuid.UUID) error
}
