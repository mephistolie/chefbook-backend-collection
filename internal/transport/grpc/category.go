package grpc

import (
	"context"
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-category/internal/entity"
	categoryFail "github.com/mephistolie/chefbook-backend-category/internal/entity/fail"
	"github.com/mephistolie/chefbook-backend-category/internal/transport/grpc/dto"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func (s *CategoryServer) GetUserCategories(_ context.Context, req *api.GetUserCategoriesRequest) (*api.GetUserCategoriesResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	categories := s.service.GetUserCategories(userId)
	if err != nil {
		return nil, err
	}

	return dto.NewGetUserCategoriesResponse(categories), nil
}

func (s *CategoryServer) GetCategoriesMap(_ context.Context, req *api.GetCategoriesMapRequest) (*api.GetCategoriesMapResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	var categoryIds []uuid.UUID
	for _, rawId := range req.CategoryIds {
		if categoryId, err := uuid.Parse(rawId); err == nil {
			categoryIds = append(categoryIds, categoryId)
		}
	}

	categories := s.service.GetCategoriesMap(categoryIds, userId)
	if err != nil {
		return nil, err
	}

	return dto.NewGetCategoriesMapResponse(categories), nil
}

func (s *CategoryServer) AddCategory(_ context.Context, req *api.AddCategoryRequest) (*api.AddCategoryResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	var categoryId uuid.UUID
	if parsedId, err := uuid.Parse(req.CategoryId); err == nil {
		categoryId = parsedId
	} else {
		categoryId = uuid.New()
	}

	if (len(req.Name)) > 64 {
		return nil, categoryFail.GrpcNameLength
	}
	if (len(req.Emoji)) > 64 {
		return nil, categoryFail.GrpcEmojiLength
	}

	var emoji *string
	if len(req.Emoji) > 0 {
		emoji = &req.Emoji
	}

	category := entity.Category{
		Id:    categoryId,
		Name:  req.Name,
		Emoji: emoji,
	}

	id, err := s.service.AddCategory(category, userId)
	if err != nil {
		return nil, err
	}

	return &api.AddCategoryResponse{CategoryId: id.String()}, nil
}

func (s *CategoryServer) GetCategory(_ context.Context, req *api.GetCategoryRequest) (*api.GetCategoryResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	categoryId, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	category, err := s.service.GetCategory(categoryId, userId)
	if err != nil {
		return nil, err
	}

	emoji := ""
	if category.Emoji != nil {
		emoji = *category.Emoji
	}

	res := api.Category{
		CategoryId: category.Id.String(),
		Name:       category.Name,
		Emoji:      emoji,
	}

	return &api.GetCategoryResponse{Category: &res}, nil
}

func (s *CategoryServer) UpdateCategory(_ context.Context, req *api.UpdateCategoryRequest) (*api.UpdateCategoryResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	categoryId, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if (len(req.Name)) > 64 {
		return nil, categoryFail.GrpcNameLength
	}
	if (len(req.Emoji)) > 64 {
		return nil, categoryFail.GrpcEmojiLength
	}

	var emoji *string
	if len(req.Emoji) > 0 {
		emoji = &req.Emoji
	}

	category := entity.Category{
		Id:    categoryId,
		Name:  req.Name,
		Emoji: emoji,
	}

	if err = s.service.UpdateCategory(category, userId); err != nil {
		return nil, err
	}

	return &api.UpdateCategoryResponse{Message: "category updated"}, nil
}

func (s *CategoryServer) DeleteCategory(_ context.Context, req *api.DeleteCategoryRequest) (*api.DeleteCategoryResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	categoryId, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.DeleteCategory(categoryId, userId); err != nil {
		return nil, err
	}

	return &api.DeleteCategoryResponse{Message: "category deleted"}, nil
}
