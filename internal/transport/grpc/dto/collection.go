package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-collection/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-collection/internal/entity"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

const (
	maxCollectionNameLength  = 64
	maxCollectionEmojiLength = 25
)

func ParseCreateCollectionRequest(req *api.CreateCollectionRequest) (entity.Collection, error) {
	if len(req.Name) == 0 {
		return entity.Collection{}, fail.GrpcInvalidBody
	}

	var collectionId *uuid.UUID
	if req.CollectionId != nil {
		if parsedId, err := uuid.Parse(*req.CollectionId); err == nil {
			collectionId = &parsedId
		}
	}
	if collectionId == nil {
		id := uuid.New()
		collectionId = &id
	}

	if len([]rune(req.Name)) > maxCollectionNameLength {
		req.Name = string([]rune(req.Name)[0:maxCollectionNameLength])
	}
	if req.Emoji != nil && len(*req.Emoji) > maxCollectionEmojiLength {
		emoji := (*req.Emoji)[0:maxCollectionEmojiLength]
		req.Emoji = &emoji
	}

	return entity.Collection{
		Id:    *collectionId,
		Name:  req.Name,
		Emoji: req.Emoji,
	}, nil
}

func ParseUpdateCollectionRequest(req *api.UpdateCollectionRequest) (entity.Collection, error) {
	if len(req.Name) == 0 {
		return entity.Collection{}, fail.GrpcInvalidBody
	}

	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return entity.Collection{}, fail.GrpcInvalidBody
	}

	if len([]rune(req.Name)) > maxCollectionNameLength {
		req.Name = string([]rune(req.Name)[0:maxCollectionNameLength])
	}
	if req.Emoji != nil && len(*req.Emoji) > maxCollectionEmojiLength {
		emoji := (*req.Emoji)[0:maxCollectionEmojiLength]
		req.Emoji = &emoji
	}

	return entity.Collection{
		Id:    collectionId,
		Name:  req.Name,
		Emoji: req.Emoji,
	}, nil
}

func NewGetUserCollectionsResponse(collections []entity.Collection) *api.GetUserCollectionsResponse {
	dtos := make([]*api.Collection, len(collections))
	for i, collection := range collections {
		dto := api.Collection{
			CollectionId: collection.Id.String(),
			Name:         collection.Name,
			Emoji:        collection.Emoji,
		}
		dtos[i] = &dto
	}
	return &api.GetUserCollectionsResponse{Collections: dtos}
}

func NewGetCollectionsMapResponse(collections map[uuid.UUID]entity.Collection) *api.GetCollectionsMapResponse {
	dtos := make(map[string]*api.Collection)
	for id, collection := range collections {
		dto := api.Collection{
			CollectionId: id.String(),
			Name:         collection.Name,
			Emoji:        collection.Emoji,
		}
		dtos[id.String()] = &dto
	}
	return &api.GetCollectionsMapResponse{Collections: dtos}
}
