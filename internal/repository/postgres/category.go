package postgres

import (
	"context"
	"log"

	"github.com/backendmagang/project-1/internal/repository/postgres/queries"
	"github.com/backendmagang/project-1/models/entity"
	"github.com/backendmagang/project-1/models/lib"
	"github.com/lib/pq"
)

func (r *repository) GetCategoryTree(ctx context.Context, req entity.GetCategoriesRequest) ([]entity.Category, error) {
	var (
		categories []entity.Category
	)

	rows, err := r.db.Query(queries.GET_CATEGORY_TREE, req.Limit, req.Offset)
	if err != nil {
		log.Println("[Repository][Postgres][GetCategories] failed to query categories, err: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.ID, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			log.Println("[Repository][Postgres][GetCategories] failed to scan category, err: ", err)
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *repository) GetCategoryDetails(ctx context.Context, req entity.GetCategoryDetailsRequest) (entity.Category, error) {
	var (
		category entity.Category
		err      error
	)

	err = r.db.QueryRow(queries.GET_CATEGORY_DETAILS, req.ID).Scan(&category.ID, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.Println("[Repository][Postgres][GetCategoryDetails] failed to scan category, err: ", err)
		return entity.Category{}, err
	}

	return category, nil
}

func (r *repository) GetCategoriesByIDs(ctx context.Context, categoryIDs []int) ([]entity.CategoryResponse, error) {
	rows, err := r.db.Query(queries.GET_CATEGORY_BY_IDS, pq.Array(categoryIDs))
	if err != nil {
		log.Println("[Repository][Postgres][GetCategoriesByIDs] failed to get categories by ids, err: ", err)
		return nil, err
	}
	defer rows.Close()

	var categories []entity.CategoryResponse

	for rows.Next() {
		var category entity.CategoryResponse
		err := rows.Scan(&category.ID, &category.Title, &category.Slug)
		if err != nil {
			log.Println("[Repository][Postgres][GetCategoriesByIDs] failed to get scan category, err: ", err)
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *repository) InsertCategory(ctx context.Context, category entity.InsertCategoryRequest) (entity.InsertCategoryRequest, error) {
	err := r.db.QueryRow(queries.INSERT_CATEGORY, category.Title, category.Slug, category.CreatedAt, category.UpdatedAt).Scan(&category.ID)
	if err != nil {
		log.Println("[Repository][Postgres][InsertCategory] failed to insert category, err: ", err)
		return entity.InsertCategoryRequest{}, err
	}

	return category, nil
}

func (r *repository) UpdateCategory(ctx context.Context, category entity.UpdateCategoryRequest) error {
	result, err := r.db.Exec(queries.UPDATE_CATEGORY, &category.Title, &category.Slug, &category.UpdatedAt, &category.ID)
	if err != nil {
		log.Println("[Repository][Postgres][UpdateCategory] failed to update category, err: ", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return lib.ErrorNoRowsAffected
	}

	return nil
}

func (r *repository) DeleteCategory(ctx context.Context, req entity.DeleteCategoryRequest) error {
	result, err := r.db.Exec(queries.DELETE_CATEGORY, req.ID)
	if err != nil {
		log.Println("[Repository][Postgres][DeleteCategory] failed to delete category, err:", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return lib.ErrorNoRowsAffected
	}

	return nil
}
