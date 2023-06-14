package service

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-category/internal/entity"
	"github.com/mephistolie/chefbook-backend-category/internal/service/category"
	"github.com/mephistolie/chefbook-backend-category/internal/service/dependencies/repository"
	"github.com/mephistolie/chefbook-backend-category/internal/service/mq"
)

type Service struct {
	Category
	MQ
}

type Category interface {
	GetUserCategories(userId uuid.UUID) []entity.Category
	GetCategoriesMap(categoryIds []uuid.UUID, userId uuid.UUID) map[uuid.UUID]entity.Category
	AddCategory(category entity.Category, userId uuid.UUID) (uuid.UUID, error)
	UpdateCategory(category entity.Category, userId uuid.UUID) error
	DeleteCategory(categoryId, userId uuid.UUID) error
}

type MQ interface {
	DeleteUser(userId uuid.UUID, messageId uuid.UUID) error
}

func New(
	repo repository.Category,
) *Service {
	return &Service{
		Category: category.NewService(repo),
		MQ:       mq.NewService(repo),
	}
}
