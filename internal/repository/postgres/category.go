package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-category/internal/entity"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func (r *Repository) GetUserCategories(userId uuid.UUID) []entity.Category {
	var categories []entity.Category

	query := fmt.Sprintf(`
		SELECT category_id, name, emoji
		FROM %s
		WHERE user_id=$1
	`, categoriesTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		log.Errorf("unable to get user %s categories: %s", userId, err)
		return []entity.Category{}
	}

	for rows.Next() {
		var category entity.Category
		err = rows.Scan(&category.Id, &category.Name, &category.Emoji)
		if err != nil {
			log.Errorf("unable to parse user %s category: %s", userId, err)
			continue
		}
		categories = append(categories, category)
	}

	return categories
}

func (r *Repository) GetCategoriesMap(categoryIds []uuid.UUID, userId uuid.UUID) map[uuid.UUID]entity.Category {
	categories := make(map[uuid.UUID]entity.Category)

	query := fmt.Sprintf(`
		SELECT category_id, name, emoji
		FROM %s
		WHERE user_id=$1 AND category_id=ANY($2)
	`, categoriesTable)

	rows, err := r.db.Query(query, userId, categoryIds)
	if err != nil {
		log.Errorf("unable to get user %s categories: %s", userId, err)
		return map[uuid.UUID]entity.Category{}
	}

	for rows.Next() {
		var category entity.Category
		err = rows.Scan(&category.Id, &category.Name, &category.Emoji)
		if err != nil {
			log.Errorf("unable to parse user %s category: %s", userId, err)
			continue
		}
		categories[category.Id] = category
	}

	return categories
}

func (r *Repository) CreateCategory(category entity.Category, userId uuid.UUID) (uuid.UUID, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (category_id, user_id, name, emoji)
		VALUES ($1, $2, $3, $4)
		RETURNING category_id
	`, categoriesTable)

	if _, err := r.db.Exec(query, category.Id, userId, category.Name, category.Emoji); err != nil {
		log.Errorf("unable to add category %s: %s", category.Id, err)
		return uuid.UUID{}, fail.GrpcInvalidBody
	}

	return category.Id, nil
}

func (r *Repository) GetCategory(categoryId uuid.UUID) (entity.Category, uuid.UUID, error) {
	var category entity.Category
	var userId uuid.UUID

	query := fmt.Sprintf(`
		SELECT category_id, user_id, name, emoji
		FROM %s
		WHERE category_id=$1
	`, categoriesTable)

	row := r.db.QueryRow(query, categoryId)
	if err := row.Scan(&category.Id, &userId, &category.Name, &category.Emoji); err != nil {
		log.Errorf("unable to get category %s: %s", category, err)
		return entity.Category{}, uuid.UUID{}, fail.GrpcNotFound
	}

	return category, userId, nil
}

func (r *Repository) UpdateCategory(category entity.Category) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET name=$1, emoji=$2
		WHERE category_id=$3
	`, categoriesTable)

	if _, err := r.db.Exec(query, category.Name, category.Emoji, category.Id); err != nil {
		log.Errorf("unable to update category %s: %s", category.Id, err)
		return fail.GrpcUnknown
	}

	return nil
}

func (r *Repository) DeleteCategory(categoryId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE category_id=$1 AND user_id=$2
	`, categoriesTable)

	if _, err := r.db.Exec(query, categoryId, userId); err != nil {
		log.Errorf("unable to delete category %s: %s", categoryId, err)
		return fail.GrpcUnknown
	}

	return nil
}
