package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-category/internal/entity"
)

func NewGetUserCategoriesResponse(categories []entity.Category) *api.GetUserCategoriesResponse {
	dtos := make([]*api.Category, len(categories))
	for i, category := range categories {
		emoji := ""
		if category.Emoji != nil {
			emoji = *category.Emoji
		}

		dto := api.Category{
			CategoryId: category.Id.String(),
			Name:       category.Name,
			Emoji:      emoji,
		}
		dtos[i] = &dto
	}
	return &api.GetUserCategoriesResponse{Categories: dtos}
}

func NewGetCategoriesMapResponse(categories map[uuid.UUID]entity.Category) *api.GetCategoriesMapResponse {
	dtos := make(map[string]*api.Category)
	for id, category := range categories {
		emoji := ""
		if category.Emoji != nil {
			emoji = *category.Emoji
		}

		dto := api.Category{
			CategoryId: id.String(),
			Name:       category.Name,
			Emoji:      emoji,
		}
		dtos[id.String()] = &dto
	}
	return &api.GetCategoriesMapResponse{Categories: dtos}
}
