package repository

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-collection/internal/entity"
)

type Collection interface {
	DeleteUser(userId, messageId uuid.UUID) error

	GetUserCollections(userId uuid.UUID) []entity.Collection
	GetCollectionsMap(collectionIds []uuid.UUID, userId uuid.UUID) map[uuid.UUID]entity.Collection
	CreateCollection(collection entity.Collection) (uuid.UUID, error)
	GetCollection(collectionId uuid.UUID) (entity.Collection, error)
	UpdateCollection(collection entity.Collection) error
	DeleteCollection(collectionId, userId uuid.UUID) error
}
