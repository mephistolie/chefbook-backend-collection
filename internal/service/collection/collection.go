package collection

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-collection/internal/entity"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func (s *Service) GetUserCollections(userId uuid.UUID) []entity.Collection {
	return s.repo.GetUserCollections(userId)
}

func (s *Service) GetCollectionsMap(collectionIds []uuid.UUID, userId uuid.UUID) map[uuid.UUID]entity.Collection {
	return s.repo.GetCollectionsMap(collectionIds, userId)
}

func (s *Service) CreateCollection(collection entity.Collection, userId uuid.UUID) (uuid.UUID, error) {
	return s.repo.CreateCollection(collection, userId)
}

func (s *Service) GetCollection(collectionId uuid.UUID, userId uuid.UUID) (entity.Collection, error) {
	collection, ownerId, err := s.repo.GetCollection(collectionId)
	if err != nil {
		return entity.Collection{}, err
	}

	if userId != ownerId {
		return entity.Collection{}, fail.GrpcAccessDenied
	}

	return collection, nil
}
func (s *Service) UpdateCollection(collection entity.Collection, userId uuid.UUID) error {
	_, ownerId, err := s.repo.GetCollection(collection.Id)
	if err != nil {
		return err
	}
	if ownerId != userId {
		return fail.GrpcAccessDenied
	}
	return s.repo.UpdateCollection(collection)
}

func (s *Service) DeleteCollection(collectionId, userId uuid.UUID) error {
	return s.repo.DeleteCollection(collectionId, userId)
}
