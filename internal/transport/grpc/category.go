package grpc

import (
	"context"
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
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

func (s *CategoryServer) CreateCategory(_ context.Context, req *api.CreateCategoryRequest) (*api.CreateCategoryResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	input, err := dto.ParseCreateCategoryRequest(req)
	if err != nil {
		return nil, err
	}

	id, err := s.service.CreateCategory(input, userId)
	if err != nil {
		return nil, err
	}

	return &api.CreateCategoryResponse{CategoryId: id.String()}, nil
}

func (s *CategoryServer) GetCategory(_ context.Context, req *api.GetCategoryRequest) (*api.Category, error) {
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

	return &api.Category{
		CategoryId: category.Id.String(),
		Name:       category.Name,
		Emoji:      category.Emoji,
	}, nil
}

func (s *CategoryServer) UpdateCategory(_ context.Context, req *api.UpdateCategoryRequest) (*api.UpdateCategoryResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	input, err := dto.ParseUpdateCategoryRequest(req)
	if err != nil {
		return nil, err
	}

	if err = s.service.UpdateCategory(input, userId); err != nil {
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
