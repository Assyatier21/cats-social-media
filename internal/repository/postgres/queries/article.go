package queries

const (
	GET_ARTICLES = `SELECT a.id, a.title, a.slug, a.html_content, a.metadata, a.created_at, a.updated_at, json_agg(c) AS categories FROM articles a
						JOIN categories c ON c.id = ANY(a.category_id)
						GROUP BY a.id LIMIT $1 OFFSET $2`

	GET_ARTICLE_DETAILS = `SELECT a.id, a.title, a.slug, a.html_content, a.metadata, a.created_at, a.updated_at,
								ARRAY_AGG(json_build_object('id', c.id, 'title', c.title, 'slug', c.slug)) as categories
								FROM articles a JOIN categories c ON c.id = ANY(a.category_id)
								WHERE a.id = $1 GROUP BY a.id`

	INSERT_ARTICLE = `INSERT INTO articles (id, title, slug, html_content, category_id, metadata, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	UPDATE_ARTICLE = `UPDATE articles SET title = $1, slug = $2, html_content = $3, category_id = $4, metadata = $5, created_at = $6, updated_at = $7
						WHERE id = $8`

	DELETE_ARTICLE = `DELETE FROM articles 
						WHERE id = $1`
)
