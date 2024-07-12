package grpc

import (
	"context"
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-collection/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-collection/internal/transport/grpc/dto"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func (s *CollectionServer) GetUserCollections(_ context.Context, req *api.GetUserCollectionsRequest) (*api.GetUserCollectionsResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	collections := s.service.GetUserCollections(userId)
	if err != nil {
		return nil, err
	}

	return dto.NewGetUserCollectionsResponse(collections), nil
}

func (s *CollectionServer) GetCollectionsMap(_ context.Context, req *api.GetCollectionsMapRequest) (*api.GetCollectionsMapResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	var collectionIds []uuid.UUID
	for _, rawId := range req.CollectionIds {
		if collectionId, err := uuid.Parse(rawId); err == nil {
			collectionIds = append(collectionIds, collectionId)
		}
	}

	collections := s.service.GetCollectionsMap(collectionIds, userId)
	if err != nil {
		return nil, err
	}

	return dto.NewGetCollectionsMapResponse(collections), nil
}

func (s *CollectionServer) CreateCollection(_ context.Context, req *api.CreateCollectionRequest) (*api.CreateCollectionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	input, err := dto.ParseCreateCollectionRequest(req)
	if err != nil {
		return nil, err
	}

	id, err := s.service.CreateCollection(input, userId)
	if err != nil {
		return nil, err
	}

	return &api.CreateCollectionResponse{CollectionId: id.String()}, nil
}

func (s *CollectionServer) GetCollection(_ context.Context, req *api.GetCollectionRequest) (*api.Collection, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	collection, err := s.service.GetCollection(collectionId, userId)
	if err != nil {
		return nil, err
	}

	return &api.Collection{
		CollectionId: collection.Id.String(),
		Name:         collection.Name,
		Emoji:        collection.Emoji,
	}, nil
}

func (s *CollectionServer) UpdateCollection(_ context.Context, req *api.UpdateCollectionRequest) (*api.UpdateCollectionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	input, err := dto.ParseUpdateCollectionRequest(req)
	if err != nil {
		return nil, err
	}

	if err = s.service.UpdateCollection(input, userId); err != nil {
		return nil, err
	}

	return &api.UpdateCollectionResponse{Message: "collection updated"}, nil
}

func (s *CollectionServer) DeleteCollection(_ context.Context, req *api.DeleteCollectionRequest) (*api.DeleteCollectionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.DeleteCollection(collectionId, userId); err != nil {
		return nil, err
	}

	return &api.DeleteCollectionResponse{Message: "collection deleted"}, nil
}
