package service

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-collection/internal/entity"
	"github.com/mephistolie/chefbook-backend-collection/internal/service/collection"
	"github.com/mephistolie/chefbook-backend-collection/internal/service/dependencies/repository"
	"github.com/mephistolie/chefbook-backend-collection/internal/service/mq"
)

type Service struct {
	Collection
	MQ
}

type Collection interface {
	GetUserCollections(userId uuid.UUID) []entity.Collection
	GetCollectionsMap(collectionIds []uuid.UUID, userId uuid.UUID) map[uuid.UUID]entity.Collection
	CreateCollection(collection entity.Collection, userId uuid.UUID) (uuid.UUID, error)
	GetCollection(collectionId uuid.UUID, userId uuid.UUID) (entity.Collection, error)
	UpdateCollection(collection entity.Collection, userId uuid.UUID) error
	DeleteCollection(collectionId, userId uuid.UUID) error
}

type MQ interface {
	DeleteUser(userId uuid.UUID, messageId uuid.UUID) error
}

func New(
	repo repository.Collection,
) *Service {
	return &Service{
		Collection: collection.NewService(repo),
		MQ:         mq.NewService(repo),
	}
}
