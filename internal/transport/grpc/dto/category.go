package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-category/internal/entity"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

const (
	maxCategoryNameLength  = 64
	maxCategoryEmojiLength = 25
)

func ParseCreateCategoryRequest(req *api.CreateCategoryRequest) (entity.Category, error) {
	if len(req.Name) == 0 {
		return entity.Category{}, fail.GrpcInvalidBody
	}

	var categoryId *uuid.UUID
	if req.CategoryId != nil {
		if parsedId, err := uuid.Parse(*req.CategoryId); err == nil {
			categoryId = &parsedId
		}
	}
	if categoryId == nil {
		id := uuid.New()
		categoryId = &id
	}

	if len(req.Name) > maxCategoryNameLength {
		req.Name = req.Name[0:maxCategoryNameLength]
	}
	if req.Emoji != nil && len(*req.Emoji) > maxCategoryEmojiLength {
		emoji := (*req.Emoji)[0:maxCategoryEmojiLength]
		req.Emoji = &emoji
	}

	return entity.Category{
		Id:    *categoryId,
		Name:  req.Name,
		Emoji: req.Emoji,
	}, nil
}

func ParseUpdateCategoryRequest(req *api.UpdateCategoryRequest) (entity.Category, error) {
	if len(req.Name) == 0 {
		return entity.Category{}, fail.GrpcInvalidBody
	}

	categoryId, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return entity.Category{}, fail.GrpcInvalidBody
	}

	if len(req.Name) > maxCategoryNameLength {
		req.Name = req.Name[0:maxCategoryNameLength]
	}
	if req.Emoji != nil && len(*req.Emoji) > maxCategoryEmojiLength {
		emoji := (*req.Emoji)[0:maxCategoryEmojiLength]
		req.Emoji = &emoji
	}

	return entity.Category{
		Id:    categoryId,
		Name:  req.Name,
		Emoji: req.Emoji,
	}, nil
}

func NewGetUserCategoriesResponse(categories []entity.Category) *api.GetUserCategoriesResponse {
	dtos := make([]*api.Category, len(categories))
	for i, category := range categories {
		dto := api.Category{
			CategoryId: category.Id.String(),
			Name:       category.Name,
			Emoji:      category.Emoji,
		}
		dtos[i] = &dto
	}
	return &api.GetUserCategoriesResponse{Categories: dtos}
}

func NewGetCategoriesMapResponse(categories map[uuid.UUID]entity.Category) *api.GetCategoriesMapResponse {
	dtos := make(map[string]*api.Category)
	for id, category := range categories {
		dto := api.Category{
			CategoryId: id.String(),
			Name:       category.Name,
			Emoji:      category.Emoji,
		}
		dtos[id.String()] = &dto
	}
	return &api.GetCategoriesMapResponse{Categories: dtos}
}
