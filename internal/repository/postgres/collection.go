package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-collection/api/model"
	"github.com/mephistolie/chefbook-backend-collection/internal/entity"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func (r *Repository) GetUserCollections(userId uuid.UUID) []entity.Collection {
	var collections []entity.Collection

	query := fmt.Sprintf(`
		SELECT collection_id, owner_id, coauthors, name, emoji
		FROM %[1]v
		LEFT JOIN
			%[2]v ON %[2]v.collection_id=%[1]v.collection_id AND %[2]v.user_id=$1
		WHERE
			%[2]v.user_id=$1 AND (%[1]v.owner_id=$1 OR %[1]v.visibility<>'%[3]v')
	`, collectionsTable, collectionsUsersTable, model.VisibilityPrivate)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		log.Errorf("unable to get user %s collections: %s", userId, err)
		return []entity.Collection{}
	}

	for rows.Next() {
		var collection entity.Collection
		err = rows.Scan(&collection.Id, &collection.OwnerId, &collection.Coauthors, &collection.Name, &collection.Emoji)
		if err != nil {
			log.Errorf("unable to parse user %s collection: %s", userId, err)
			continue
		}
		collections = append(collections, collection)
	}

	return collections
}

func (r *Repository) GetCollectionsMap(collectionIds []uuid.UUID, userId uuid.UUID) map[uuid.UUID]entity.Collection {
	collections := make(map[uuid.UUID]entity.Collection)

	query := fmt.Sprintf(`
		SELECT collection_id, owner_id, coauthors, name, emoji
		FROM %[1]v
		LEFT JOIN
			%[2]v ON %[2]v.collection_id=%[1]v.collection_id AND %[2]v.user_id=$1
		WHERE
			%[2]v.user_id=$1 AND collection_id=ANY($2) AND (%[1]v.owner_id=$1 OR %[1]v.visibility<>'%[3]v')
	`, collectionsTable, collectionsUsersTable, model.VisibilityPrivate)

	rows, err := r.db.Query(query, userId, collectionIds)
	if err != nil {
		log.Errorf("unable to get user %s collections: %s", userId, err)
		return map[uuid.UUID]entity.Collection{}
	}

	for rows.Next() {
		var collection entity.Collection
		err = rows.Scan(&collection.Id, &collection.OwnerId, &collection.Coauthors, &collection.Name, &collection.Emoji)
		if err != nil {
			log.Errorf("unable to parse user %s collection: %s", userId, err)
			continue
		}
		collections[collection.Id] = collection
	}

	return collections
}

func (r *Repository) CreateCollection(collection entity.Collection) (uuid.UUID, error) {
	tx, err := r.startTransaction()
	if err != nil {
		return uuid.UUID{}, err
	}

	createCollectionQuery := fmt.Sprintf(`
		INSERT INTO %s (collection_id, owner_id, coauthors, name, emoji)
		VALUES ($1, $2, $3, $4)
		RETURNING collection_id
	`, collectionsTable)

	if _, err = tx.Exec(createCollectionQuery, collection.Id, collection.OwnerId, collection.Name, collection.Emoji); err != nil {
		log.Errorf("unable to add collection %s: %s", collection.Id, err)
		return uuid.UUID{}, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	addCollectionUserQuery := fmt.Sprintf(`
			INSERT INTO %s (collection_id, user_id)
			VALUES ($1, $2)
		`, collectionsUsersTable)

	if _, err = tx.Exec(addCollectionUserQuery, collection.Id, userId); err != nil {
		log.Errorf("unable to add user for collection %s: %s", collection.Id, err)
		return uuid.UUID{}, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return collection.Id, nil
}

func (r *Repository) GetCollection(collectionId uuid.UUID) (entity.Collection, error) {
	var collection entity.Collection

	query := fmt.Sprintf(`
		SELECT collection_id, owner_id, name, emoji
		FROM %s
		WHERE collection_id=$1
	`, collectionsTable)

	row := r.db.QueryRow(query, collectionId)
	if err := row.Scan(&collection.Id, &collection.OwnerId, &collection.Name, &collection.Emoji); err != nil {
		log.Errorf("unable to get collection %s: %s", collectionId, err)
		return entity.Collection{}, fail.GrpcNotFound
	}

	return collection, nil
}

func (r *Repository) UpdateCollection(collection entity.Collection) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET name=$1, emoji=$2
		WHERE collection_id=$3 AND 
	`, collectionsTable)

	if _, err := r.db.Exec(query, collection.Name, collection.Emoji, collection.Id); err != nil {
		log.Errorf("unable to update collection %s: %s", collection.Id, err)
		return fail.GrpcUnknown
	}

	return nil
}

func (r *Repository) DeleteCollection(collectionId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE collection_id=$1 AND user_id=$2
	`, collectionsTable)

	if _, err := r.db.Exec(query, collectionId, userId); err != nil {
		log.Errorf("unable to delete collection %s: %s", collectionId, err)
		return fail.GrpcUnknown
	}

	return nil
}
