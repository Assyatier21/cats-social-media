package queries

const (
	GET_CATEGORY_TREE = `SELECT id, title, slug, created_at, updated_at
                     FROM categories ORDER BY id LIMIT $1 OFFSET $2`

	GET_CATEGORY_DETAILS = `SELECT id, title, slug, created_at, updated_at
                        FROM categories WHERE id = $1`

	GET_CATEGORY_BY_IDS = `SELECT id, title, slug FROM categories WHERE id IN (
  								SELECT unnest($1::integer[]));`

	INSERT_CATEGORY = `INSERT INTO categories (title, slug, created_at, updated_at)
							VALUES ($1, $2, $3, $4) RETURNING id`

	UPDATE_CATEGORY = `UPDATE categories SET title = $1, slug = $2, updated_at = $3
							WHERE id = $4`

	// Query if Category Used by Article, Cancel Commit
	DELETE_CATEGORY = `	WITH used_categories AS (
						 	SELECT DISTINCT unnest(category_id) AS category
							FROM articles WHERE $1 = ANY(category_id) LIMIT 1
				 	   	)
					   	DELETE FROM categories WHERE id = $1
					   	AND NOT EXISTS (
							SELECT 1 FROM used_categories
					   	);`
)
